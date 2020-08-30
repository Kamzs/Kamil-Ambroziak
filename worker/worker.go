package worker

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/utils"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

var storageCli fetchers.Storage

func init() {
	storageCli, _ = storage.NewMySQL()

}

func NewWorker() *worker {
	c := cron.New(cron.WithSeconds())
	go c.Run()
	return &worker{
		cron: c,
	}
}

type worker struct {
	cron *cron.Cron
}

func (w *worker) RegisterFetcher(fetcher *fetchers.Fetcher) (cron.EntryID, utils.RestErr) {
	jobID, err := w.cron.AddFunc(fmt.Sprintf("*/%v * * * * *", fetcher.Interval), func() { doJob(fetcher) })
	if err != nil {
		return 0, utils.NewInternalServerError("job could not be registered by worker", err)
	}
	return jobID, nil
}
func (w *worker) DeregisterFetcher(jobID cron.EntryID) {
	w.cron.Remove(jobID)
}

func doJob(fetcher *fetchers.Fetcher) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", fetcher.Url, nil)
	if err != nil {
		fmt.Println("error request")
		return
	}
	var start, connect, dns, tlsHandshake time.Time
	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Done: %v\n", time.Since(dns))
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			fmt.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("Connect time: %v\n", time.Since(connect))
		},

		GotFirstResponseByte: func() {
			fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
		},
	}
	ctxTimer := httptrace.WithClientTrace(ctxTimeout, trace)
	historyEl := &fetchers.HistoryElement{
		Id:        fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	//res, err := http.DefaultClient.Do(req.WithContext(ctxTimer))
	start = time.Now()
	res, err := http.DefaultTransport.RoundTrip(req.WithContext(ctxTimer))

	totalResponseTime := time.Since(start)
	if err != nil {
		if fmt.Sprintf("%T", err) == "context.deadlineExceededError" {
			historyEl.Response = ""
			historyEl.Duration = 5
			_ = storageCli.SaveHistoryForFetcher(historyEl)
		}
		//todo add logger

	} else {
		if res.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
				return
			}
			historyEl.Response = string(bodyBytes)
			/*			nanoseconds := totalResponseTime - start
						//duration := time.Duration(nanosecods)
						duration := float64(nanoseconds)/float64(time.Second)*/
			/*			fmt.Println(duration)
						fmt.Printf("%.2f", duration)*/
			historyEl.Duration = totalResponseTime.Seconds()
		}
		_ = storageCli.SaveHistoryForFetcher(historyEl)
	}
}
