package fetchers

import (
	"Kamil-Ambroziak/utils"
	"strings"
)

//todo change to interface in order to allow mocks creation

type Storage interface {
	SaveFetcher(fetcher *Fetcher) utils.RestErr
	UpdateFetcher(fetcher *Fetcher) utils.RestErr
	DeleteFetcher(fetcherId int64) utils.RestErr
	FindAllFetchers() ([]Fetcher, utils.RestErr)
	GetHistoryForFetcher(fetcherID int64) (*Fetcher,utils.RestErr)
}
type Fetcher struct {
	Id       int64  `json:"id"`
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}

type Fetchers []Fetcher

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
