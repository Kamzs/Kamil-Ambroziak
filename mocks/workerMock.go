package mocks

import (
	"errors"

	fetchers "github.com/Kamzs/Kamil-Ambroziak"
	"github.com/Kamzs/Kamil-Ambroziak/utils"
	"github.com/robfig/cron/v3"
)

const exampleId = 0

type WorkerMock struct {
	RegisterFetcherError   bool
	DeregisterFetcherError bool

	ErrAfterSuccess int
	counter         int
}

func (w *WorkerMock) RegisterFetcher(fetcher *fetchers.Fetcher) (cron.EntryID, utils.RestErr) {
	if w.RegisterFetcherError {
		return exampleId, utils.NewInternalServerError(exampleErrorMsg, errors.New(exampleErrorMsg))
	}
	return exampleId, nil
}
func (w *WorkerMock) DeregisterFetcher(jobID cron.EntryID) {
}
