package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"strconv"
)

type JsonWithID struct {
	Id          int64  `json:"id"`
}
type FetcherUpdateResponse struct {
	Id       int64 `json:"id"`
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}
type GetAllFetchersResponse struct {
	Id       int64  `json:"id"`
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}
type HistoryElementResponse struct {
	Response  string  `json:"response"`
	Duration  float64 `json:"duration"`
	CreatedAt int64   `json:"created_at"`
}
func getFetcherId(fetcherIdParam string) (int64, utils.RestErr) {
	fetcherId, fetcherErr := strconv.ParseInt(fetcherIdParam, 10, 64)
	if fetcherErr != nil {
		return 0, utils.NewBadRequestError("fetcher id should be an int64")
	}
	return fetcherId, nil
}
func fillMissingFields(oldFetcher *fetchers.Fetcher, newFetcher *fetchers.Fetcher){
	if newFetcher.Interval == 0 {newFetcher.Interval = oldFetcher.Interval}
	if newFetcher.Url == "" {newFetcher.Url = oldFetcher.Url}
}
