package worker

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/utils"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)
var storageCli fetchers.Storage

func init(){
	storageCli,_ = storage.NewMySQL()

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

func doJob2(fetcher *fetchers.Fetcher) {
	historyEl := &fetchers.HistoryElement{
		Id: fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	var timeout time.Duration = 5000
	//http.DefaultClient.Timeout = timeout
	resp, err := http.Get(fetcher.Url)
	if err != nil {
		print(err)
	}
/*	defer resp.Body.Close()
	ctx := context.Background()
	context.WithTimeout(ctx, time.Second*15)
	req, _ := http.NewRequestWithContext(ctx, "get", fetcher.Url, nil)
	//client := http.Client{Timeout: time.Second*5}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}*/
//	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		historyEl.Response = string(bodyBytes)
		historyEl.Duration = timeout.Seconds()
	}
	if resp.StatusCode == http.StatusRequestTimeout{
		historyEl.Response = ""
		historyEl.Duration = timeout.Seconds()
	}

	_ = storageCli.SaveHistoryForFetcher(historyEl)

}
//not working
func doJob3(fetcher *fetchers.Fetcher) {
	req, err := http.NewRequest("GET", fetcher.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a httpstat powered context
//	ctx,_ := context.WithTimeout(context.Background(),5000)
//	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.Client{Timeout: 5000}
	start := time.Now().Unix()
	res, err := client.Do(req)
	end := time.Now().Unix()
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	historyEl := &fetchers.HistoryElement{
		Id: fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		historyEl.Response = string(bodyBytes)
		durationInt64 := end - start
		duration := time.Duration(durationInt64)
		historyEl.Duration = duration.Seconds()
	}
	if res.StatusCode == http.StatusRequestTimeout{
		historyEl.Response = ""
		historyEl.Duration = 5
	}

	_ = storageCli.SaveHistoryForFetcher(historyEl)

}
//working
func doJob (fetcher *fetchers.Fetcher){
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
	ctxTimer := httptrace.WithClientTrace(ctxTimeout,trace)
	historyEl := &fetchers.HistoryElement{
		Id: fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	//res, err := http.DefaultClient.Do(req.WithContext(ctxTimer))
	start = time.Now()
	res, err := http.DefaultTransport.RoundTrip(req.WithContext(ctxTimer))

	totalResponseTime := time.Since(start)
	if err != nil {
		fmt.Println("error")
		historyEl.Response = ""
		historyEl.Duration = 5
		_ = storageCli.SaveHistoryForFetcher(historyEl)
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

func doJob4(fetcher *fetchers.Fetcher) {

	c:= make(chan struct{})
	timer:= time.AfterFunc(5*time.Second, func() {
		close(c)
	})

	req, err := http.NewRequest("GET","http://httpbin.org/range/2048?duration=8&chunk_size=256",nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Cancel = c

	log.Println("sending request")
	resp,err := http.DefaultClient.Do(req)
	if err!= nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Println("reading body")

	for {
		//timer.Reset(2*time.Second)
		timer.Reset(50*time.Millisecond)
		_,err = io.CopyN(ioutil.Discard, resp.Body, 256)
		if err ==io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
/*
	ctx, cancel := context.WithTimeout(context.Background(),5000)
	defer cancel ()

	go func () {
		select {
		case <- time.After(10*time.Second):
			fmt.Println("overslept")
		case <- ctx.Done():
			fmt.Println(ctx.Err())
		}
	}()
	req, _ := http.NewRequestWithContext(ctx, "GET", fetcher.Url,nil)

	start := time.Now().Unix()
	res, _ := http.DefaultClient.Do(req)
	end := time.Now().Unix()

	req, err := http.NewRequest("GET", fetcher.Url, nil)
	if err != nil {
		log.Fatal(err)
	}

	historyEl := &fetchers.HistoryElement{
		Id: fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		historyEl.Response = string(bodyBytes)
		durationInt64 := end - start
		duration := time.Duration(durationInt64)
		historyEl.Duration = duration.Seconds()
	}
	if res.StatusCode == http.StatusRequestTimeout{
		historyEl.Response = ""
		historyEl.Duration = 5
	}

	_ = storageCli.SaveHistoryForFetcher(historyEl)*/

}