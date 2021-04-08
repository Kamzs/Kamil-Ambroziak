package fetchers

import (
	"net/url"

	"github.com/Kamzs/Kamil-Ambroziak/logger"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
	"github.com/robfig/cron/v3"
)

type Storage interface {
	SaveFetcher(fetcher *Fetcher) utils.RestErr
	UpdateFetcher(fetcher *Fetcher) utils.RestErr
	DeleteFetcher(fetcherId int64) utils.RestErr
	FindAllFetchers() ([]Fetcher, utils.RestErr)
	GetFetcher(fetcherId int64) (*Fetcher, utils.RestErr)
	SaveHistoryForFetcher(historyEl *HistoryElement) utils.RestErr
	GetHistoryForFetcher(fetcherId int64) ([]HistoryElement, utils.RestErr)
}

type Worker interface {
	RegisterFetcher(fetcher *Fetcher) (cron.EntryID, utils.RestErr)
	DeregisterFetcher(jobId cron.EntryID)
}

type Fetcher struct {
	Id       int64 `json:"id"`
	JobID    int64
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}

type HistoryElement struct {
	Id        int64   `json:"id"`
	Response  string  `json:"response"`
	Duration  float64 `json:"duration"`
	CreatedAt int64   `json:"created_at"`
}

func (fetcher *Fetcher) Validate(update bool) utils.RestErr {
	checkInterval := true
	checkURL := true
	if update == true {
		if fetcher.Interval == 0 {
			checkInterval = false
		}
		if fetcher.Url == "" {
			checkURL = false
		}
	}
	if checkInterval {
		if fetcher.Interval <= 0 {
			return utils.NewBadRequestError("invalid interval")
		}
	}
	if checkURL {
		_, err := url.ParseRequestURI(fetcher.Url)
		if err != nil {
			logger.Error("provided wrong url", err)
			return utils.NewBadRequestError("provided wrong url")
		}
	}
	return nil
}
