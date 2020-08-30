package worker

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/logger"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/utils"
	"context"
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
		logger.Error(fmt.Sprintf("error when trying prepare request for job %v", fetcher.JobID), err)
		return
	}
	trace := &httptrace.ClientTrace{}
	ctxTimer := httptrace.WithClientTrace(ctxTimeout, trace)

	historyEl := &fetchers.HistoryElement{
		Id:        fetcher.Id,
		CreatedAt: time.Now().Unix(),
	}
	start := time.Now()
	res, err := http.DefaultTransport.RoundTrip(req.WithContext(ctxTimer))
	totalResponseTime := time.Since(start)
	if err != nil {
		if fmt.Sprintf("%T", err) == "context.deadlineExceededError" {
			historyEl.Response = ""
			historyEl.Duration = 5
			err = storageCli.SaveHistoryForFetcher(historyEl)
			if err != nil {
				logger.Error(fmt.Sprintf("history for fetcher = fetcherId %v could not be saved", fetcher.Id), err)
			}
		}
		logger.Error(fmt.Sprintf("page %s could not be raeched", fetcher.Url), err)
	} else {
		if res.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
				return
			}
			historyEl.Response = string(bodyBytes)
			historyEl.Duration = totalResponseTime.Seconds()
		}
		err = storageCli.SaveHistoryForFetcher(historyEl)
		if err != nil {
			logger.Error(fmt.Sprintf("history for fetcher = fetcherId %v could not be saved", fetcher.Id), err)
		}
	}
}
