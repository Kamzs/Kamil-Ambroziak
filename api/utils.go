package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"strconv"
)

type JsonWithID struct {
	Id          int64  `json:"id"`
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