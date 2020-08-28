package fetchers

import (
	"Kamil-Ambroziak/utils"
	"strings"
)

//todo change to interface in order to allow mocks creation

type Fetcher struct {
	Id          int64  `json:"id"`
	Url	string `json:"url"`
	Interval int64 `json:"interval"`
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
