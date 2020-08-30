package mocks

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/utils"
	"errors"
)

const exampleErrorMsg = "error"

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

func (db *MySQLMock) SaveFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	if db.SaveFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return nil
}
func (db *MySQLMock) UpdateFetcher(fetcher *fetchers.Fetcher) utils.RestErr {
	if db.UpdateFetcherError {
		return utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
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
	return []fetchers.Fetcher{}, nil
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
