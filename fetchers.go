package fetchers

import (
	"Kamil-Ambroziak/utils"
	"github.com/robfig/cron/v3"
	"strings"
)

//todo change to interface in order to allow mocks creation

type Storage interface {
	SaveFetcher(fetcher *Fetcher) utils.RestErr
	UpdateFetcher(fetcher *Fetcher) utils.RestErr
	DeleteFetcher(fetcherId int64) utils.RestErr
	FindAllFetchers() ([]Fetcher, utils.RestErr)
	GetFetcher(fetcherId int64 ) (*Fetcher, utils.RestErr)
	GetHistoryForFetcher(fetcherId int64 ) ([]HistoryElement, utils.RestErr)
}

type Worker interface {
	RegisterFetcher(fetcher *Fetcher) (cron.EntryID, utils.RestErr)
	DeregisterFetcher(jobId cron.EntryID)
	UpdateFetcher(fetcher *Fetcher) utils.RestErr
}

type Fetcher struct {
	Id       int64 `json:"id"`
	JobID    int64
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}
type GetAllFetchersResponse struct {
	Id       int64 `json:"id"`
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}

type HistoryElement struct {
	Id        int64  `json:"id"`
	Response  string `json:"url"`
	Duration  int64  `json:"duration"`
	CreatedAt int64  `json:"created_at"`
}
type HistoryElementResponse struct {
	Response  string `json:"url"`
	Duration  int64  `json:"duration"`
	CreatedAt int64  `json:"created_at"`
}

func (fetcher *Fetcher) Validate() utils.RestErr {

	fetcher.Url = strings.TrimSpace(fetcher.Url)
	//todo add more checkings
	if fetcher.Url == "" {
		return utils.NewBadRequestError("invalid url")
	}
	//todo add more checkings
	if fetcher.Interval == 0 {
		return utils.NewBadRequestError("invalid interval")
	}
	return nil
}
