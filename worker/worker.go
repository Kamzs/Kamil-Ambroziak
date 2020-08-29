package worker


import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"fmt"
	"github.com/robfig/cron/v3"
)

func NewWorker () *worker{
	c := cron.New(cron.WithSeconds())
	go c.Run()
	return &worker{
		cron: c,
	}
}

type worker struct {
	cron *cron.Cron
}

func (c *worker) RegisterFetcher(fetcher *fetchers.Fetcher) (cron.EntryID, utils.RestErr){
	jobID,err := c.cron.AddFunc(fmt.Sprintf("*/%v * * * * *", fetcher.Interval), func() { fmt.Println("function") })
	if err!= nil {
		return 0,utils.NewInternalServerError("job could not be registered by worker", err)
	}
	return jobID,nil
}
func (c *worker)DeregisterFetcher(jobID cron.EntryID){
	c.cron.Remove(jobID)
}
func (c *worker)UpdateFetcher(fetcher *fetchers.Fetcher) utils.RestErr{
	return nil
}



//todo implement fetching data
//todo implement saving fetched data
