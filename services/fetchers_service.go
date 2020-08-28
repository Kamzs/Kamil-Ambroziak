package services


import (
	"Kamil-Ambroziak/domain/fetchers"
	"Kamil-Ambroziak/utils"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

var (
	FetchersService fetchersServiceInterface = &fetchersService{}
)

type fetchersService struct{}

type fetchersServiceInterface interface {
	CreateFetcher(fetchers.Fetcher) (*fetchers.Fetcher, utils.RestErr)
	UpdateFetcher(fetchers.Fetcher) (*fetchers.Fetcher, utils.RestErr)
	DeleteFetcher(int64) utils.RestErr
	FindAllFetchers() ([]fetchers.Fetcher, utils.RestErr)
	GetHistoryForFetcher(int64) (*fetchers.Fetcher, utils.RestErr)
}

func (s *fetchersService) CreateFetcher(fetcher fetchers.Fetcher) (*fetchers.Fetcher, utils.RestErr) {
	if err := fetcher.Validate(); err != nil {
		return nil, err
	}
	if err := fetcher.SaveFetcher(); err != nil {
		return nil, err
	}
	//todo implement
	registerWorker()
	return &fetcher, nil
}

func (s *fetchersService) UpdateFetcher(fetcher fetchers.Fetcher) (*fetchers.Fetcher, utils.RestErr) {
	//todo --change to updating without prior getting
	current := &fetchers.Fetcher{Id: fetcher.Id}
	if err := current.GetHistoryForFetcher(); err != nil {
		return nil, err
	}
		if fetcher.Url != "" {
			current.Url = fetcher.Url
		}

		if fetcher.Interval != 0 {
			current.Interval = fetcher.Interval
		}
	if err := current.UpdateFetcher(); err != nil {
		return nil, err
	}
	//todo implement
	//updateWorker()
	return current, nil
}

func (s *fetchersService) DeleteFetcher(fetcherId int64) utils.RestErr {
	dao := &fetchers.Fetcher{Id: fetcherId}
	//todo implement
	//deleteWorker()
	return dao.DeleteFetcher()
}

func (s *fetchersService) FindAllFetchers() ([]fetchers.Fetcher, utils.RestErr) {
	dao := &fetchers.Fetcher{}
	return dao.FindAllFetchers()
}
func (s *fetchersService) GetHistoryForFetcher(fetcherId int64) (*fetchers.Fetcher, utils.RestErr) {
	dao := &fetchers.Fetcher{Id: fetcherId}
	if err := dao.GetHistoryForFetcher(); err != nil {
		return nil, err
	}
	return dao, nil
}
//v0.0.1
func registerWorker(){
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() { fmt.Println("getPage1 1 sec") })

	// Start cron with one scheduled job
	c.Start()
	time.Sleep(2 * time.Minute)

	// Funcs may also be added to a running Cron
	//not working
	//	entryID2, _ := c.AddFunc("*/2 * * * *", func() { fmt.Println("getPage2 2 sec") })
	c.AddFunc("*/2 * * * *", func() { fmt.Println("getPage2 2 sec") })
	time.Sleep(5 * time.Minute)

	//Remove Job2 and add new Job2 that run every 1 minute
	//not working
	//c.Remove(entryID2)
	fmt.Println("getPage2 2 sec deleted")
	c.AddFunc("*/1 * * * *", func() { fmt.Println("getPage2 1 sec")})
	time.Sleep(5 * time.Minute)

	//todo implement fetching data
	//todo implement saving fetched data


}