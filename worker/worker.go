package worker

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/storage"
	"Kamil-Ambroziak/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net/http"
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

func doJob(fetcher *fetchers.Fetcher) {
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
	//defer resp.Body.Close()
	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}

	historyEl := &fetchers.HistoryElement{
		CreatedAt: 2,
		Id:        fetcher.Id,
		Response:  bodyString,
		Duration:  2,
	}

	_ = storageCli.SaveHistoryForFetcher(historyEl)

}
