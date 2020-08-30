package mocks

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"errors"
)

const (
	FetcherId       = 1
	exampleErrorMsg = "error"
	okUrl = "https://httpbin.org/range/10"
	badUrl = "ttpbin.org/range/10"
	entityTooLargeUrl = "https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10https://httpbin.org/range/10"
	okInterval = 10
	badInterval = -1
)

type MySQLMock struct {
	SaveFetcherError bool
	UpdateFetcherError bool
	DeleteFetcherError bool
	FindAllFetchersError bool
	GetFetcherError bool
	SaveHistoryForFetcherError bool
	GetHistoryForFetcherError bool

	Fetcher                 *fetchers.Fetcher

	ErrAfterSuccess          int
	counter                  int
}
type FetcherBadBody struct {
	Url bool
}
func GetFetcherBadBody() *FetcherBadBody{
	return &FetcherBadBody{
		Url: true,
	}
}
func GetFetcherEntityToBig() *fetchers.Fetcher{
	return &fetchers.Fetcher{
		Url: entityTooLargeUrl,
		Interval: okInterval,
	}
}
func GetFetcherOk() *fetchers.Fetcher{
	return &fetchers.Fetcher{
		Url: okUrl,
		Interval: okInterval,
	}
}
func GetFetcherIntervalError() *fetchers.Fetcher{
	return &fetchers.Fetcher{
		Url: okUrl,
		Interval: badInterval,
	}
}
func GetFetcherUrlError() *fetchers.Fetcher{
	return &fetchers.Fetcher{
		Url: badUrl,
		Interval: okInterval,
	}
}
func (db *MySQLMock) SaveFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	if db.SaveFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	fetcher.Id = FetcherId
	return nil
}
func (db *MySQLMock) UpdateFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	if db.UpdateFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	if db.ErrAfterSuccess != 0 && db.counter == db.ErrAfterSuccess {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	if db.ErrAfterSuccess > 0 {
		db.counter++
	}
	fetcher.Id = FetcherId
	return nil
}
func (db *MySQLMock) DeleteFetcher(fetcherId int64) utils.RestErr {
	if db.DeleteFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return nil
}
func (db *MySQLMock) FindAllFetchers() ([]fetchers.Fetcher, utils.RestErr) {
	if db.FindAllFetchersError {
		return nil,utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return []fetchers.Fetcher{{Id: FetcherId}}, nil
}
func (db *MySQLMock) GetFetcher(fetcherId int64) (*fetchers.Fetcher, utils.RestErr) {
	if db.GetFetcherError {
		return nil,utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return &fetchers.Fetcher{}, nil
}
func (db *MySQLMock) SaveHistoryForFetcher(historyEl *fetchers.HistoryElement) utils.RestErr {
	if db.SaveHistoryForFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return nil
}
func (db *MySQLMock) GetHistoryForFetcher(fetcherId int64 ) ([]fetchers.HistoryElement, utils.RestErr) {
	if db.GetHistoryForFetcherError {
		return nil,utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return []fetchers.HistoryElement{},nil
}
