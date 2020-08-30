package api

import (
	fetchers "Kamil-Ambroziak"
	"Kamil-Ambroziak/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


//test names
const (
	negative_entityTooLarge           = "negative_entityTooLarge"
	negative_badBody                  = "negative_badBody"
	negative_workerError              = "negative_workerError"
	negative_mysqlError               = "negative_mysqlError"
	negative_mysqlError_second        = "negative_mysqlError_second"
	negative_validationError_interval = "negative_validationError_interval"
	negative_validationError_url      = "negative_validationError_url"
	negative_badFormatId              = "negative_badFormatId"
	positive                          = "positive"
)

func TestApi_AddFetcher(t *testing.T) {
	type fields struct {
		Storage fetchers.Storage
		Worker  fetchers.Worker
	}
	type args struct {
		fetcher        *fetchers.Fetcher
		fetcherBadBody *mocks.FetcherBadBody
		wrongId        string
		IdOk           int64
		wantCode       int
		wantId         int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: positive,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				wantCode: http.StatusCreated,
				wantId:   mocks.FetcherId,
			},
		},
		{
			name: negative_validationError_interval,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherIntervalError(),
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_validationError_url,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherUrlError(),
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_mysqlError,
			fields: fields{
				Storage: &mocks.MySQLMock{SaveFetcherError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				wantCode: http.StatusInternalServerError,
			},
		},
		{
			name: negative_workerError,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{RegisterFetcherError: true},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				wantCode: http.StatusInternalServerError,
			},
		},
		{
			name: negative_badBody,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcherBadBody: mocks.GetFetcherBadBody(),
				wantCode:       http.StatusBadRequest,
			},
		},
		{
			name: negative_entityTooLarge,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherEntityToBig(),
				wantCode: http.StatusRequestEntityTooLarge,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAPIServer(tt.fields.Storage, tt.fields.Worker)
			server := httptest.NewServer(api.Router)
			url := fmt.Sprintf("%s/api/fetcher", server.URL)
			var serialized []byte
			if tt.args.fetcherBadBody != nil {
				serialized, _ = json.Marshal(tt.args.fetcherBadBody)
			} else {
				serialized, _ = json.Marshal(tt.args.fetcher)
			}
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(serialized))
			resp := execReq(req, t)
			respBody := &JsonWithID{}
			_ = json.NewDecoder(resp.Body).Decode(&respBody)
			if tt.args.wantCode != 0 {
				checkRespCode(t, tt.args.wantCode, resp.StatusCode)
			}
			if tt.args.wantId != 0 {
				checkIntInBody(t, tt.args.wantId, respBody.Id)
			}
		})
	}
}

func TestApi_GetAllFetchers(t *testing.T) {
	type fields struct {
		Storage fetchers.Storage
		Worker  fetchers.Worker
	}
	type args struct {
		fetcher        *fetchers.Fetcher
		fetcherBadBody *mocks.FetcherBadBody
		wrongId        string
		IdOk           int64
		wantCode       int
		wantId         int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: positive,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wantCode: http.StatusOK,
				wantId:   mocks.FetcherId,
			},
		},
		{
			name: negative_mysqlError,
			fields: fields{
				Storage: &mocks.MySQLMock{FindAllFetchersError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				wantCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAPIServer(tt.fields.Storage, tt.fields.Worker)
			server := httptest.NewServer(api.Router)
			url := fmt.Sprintf("%s/api/fetcher", server.URL)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			resp := execReq(req, t)
			var respBody []GetAllFetchersResponse
			_ = json.NewDecoder(resp.Body).Decode(&respBody)
			if tt.args.wantCode != 0 {
				checkRespCode(t, tt.args.wantCode, resp.StatusCode)
			}
			if tt.args.wantId != 0 {
				checkIntInBody(t, tt.args.wantId, respBody[0].Id)
			}
		})
	}
}

func TestApi_UpdateFetcher(t *testing.T) {
	type fields struct {
		Storage fetchers.Storage
		Worker  fetchers.Worker
	}
	type args struct {
		fetcher        *fetchers.Fetcher
		fetcherBadBody *mocks.FetcherBadBody
		wrongId        string
		IdOk           int64
		wantCode       int
		wantId         int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: positive,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				wantCode: http.StatusOK,
				IdOk:     mocks.FetcherId,
				wantId:   mocks.FetcherId,
			},
		},
		{
			name: negative_validationError_interval,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherIntervalError(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_badFormatId,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wrongId:  mocks.WrongId,
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_validationError_url,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherUrlError(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_mysqlError,
			fields: fields{
				Storage: &mocks.MySQLMock{GetFetcherError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusInternalServerError,
			},
		},
		{
			name: negative_mysqlError_second,
			fields: fields{
				Storage: &mocks.MySQLMock{UpdateFetcherError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusInternalServerError,
			},
		},
		{
			name: negative_workerError,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{RegisterFetcherError: true},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusInternalServerError,
			},
		},
		{
			name: negative_badBody,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcherBadBody: mocks.GetFetcherBadBody(),
				IdOk:           mocks.FetcherId,
				wantCode:       http.StatusBadRequest,
			},
		},
		{
			name: negative_entityTooLarge,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherEntityToBig(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusRequestEntityTooLarge,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAPIServer(tt.fields.Storage, tt.fields.Worker)
			server := httptest.NewServer(api.Router)
			var url string
			if tt.args.IdOk != 0 {
				url = fmt.Sprintf("%s/api/fetcher/%v", server.URL, tt.args.IdOk)
			} else {
				url = fmt.Sprintf("%s/api/fetcher/%s", server.URL, tt.args.wrongId)
			}
			var serialized []byte
			if tt.args.fetcherBadBody != nil {
				serialized, _ = json.Marshal(tt.args.fetcherBadBody)
			} else {
				serialized, _ = json.Marshal(tt.args.fetcher)
			}
			req, _ := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(serialized))
			resp := execReq(req, t)
			respBody := &FetcherUpdateResponse{}
			_ = json.NewDecoder(resp.Body).Decode(&respBody)
			if tt.args.wantCode != 0 {
				checkRespCode(t, tt.args.wantCode, resp.StatusCode)
			}
			if tt.args.wantId != 0 {
				checkIntInBody(t, tt.args.wantId, respBody.Id)
			}
		})
	}
}

func TestApi_DeleteFetcher(t *testing.T) {
	type fields struct {
		Storage fetchers.Storage
		Worker  fetchers.Worker
	}
	type args struct {
		fetcher        *fetchers.Fetcher
		fetcherBadBody *mocks.FetcherBadBody
		wrongId        string
		IdOk           int64
		wantCode       int
		wantId         int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: positive,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wantCode: http.StatusOK,
				IdOk:     mocks.FetcherId,
				wantId:   mocks.FetcherId,
			},
		},
		{
			name: negative_badFormatId,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wrongId:  mocks.WrongId,
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_mysqlError,
			fields: fields{
				Storage: &mocks.MySQLMock{DeleteFetcherError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAPIServer(tt.fields.Storage, tt.fields.Worker)
			server := httptest.NewServer(api.Router)
			var url string
			if tt.args.IdOk != 0 {
				url = fmt.Sprintf("%s/api/fetcher/%v", server.URL, tt.args.IdOk)
			} else {
				url = fmt.Sprintf("%s/api/fetcher/%s", server.URL, tt.args.wrongId)
			}
			req, _ := http.NewRequest(http.MethodDelete, url, nil)
			resp := execReq(req, t)
			respBody := &JsonWithID{}
			_ = json.NewDecoder(resp.Body).Decode(&respBody)
			if tt.args.wantCode != 0 {
				checkRespCode(t, tt.args.wantCode, resp.StatusCode)
			}
			if tt.args.wantId != 0 {
				checkIntInBody(t, tt.args.wantId, respBody.Id)
			}
		})
	}
}

func TestApi_GetHistoryForFetcher(t *testing.T) {
	type fields struct {
		Storage fetchers.Storage
		Worker  fetchers.Worker
	}
	type args struct {
		fetcher        *fetchers.Fetcher
		fetcherBadBody *mocks.FetcherBadBody
		wrongId        string
		IdOk           int64
		wantCode       int
		wantCreatedAt  int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: positive,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wantCode:      http.StatusOK,
				IdOk:          mocks.FetcherId,
				wantCreatedAt: mocks.CreatedAt,
			},
		},
		{
			name: negative_badFormatId,
			fields: fields{
				Storage: &mocks.MySQLMock{},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				wrongId:  mocks.WrongId,
				wantCode: http.StatusBadRequest,
			},
		},
		{
			name: negative_mysqlError,
			fields: fields{
				Storage: &mocks.MySQLMock{GetHistoryForFetcherError: true},
				Worker:  &mocks.WorkerMock{},
			},
			args: args{
				fetcher:  mocks.GetFetcherOk(),
				IdOk:     mocks.FetcherId,
				wantCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAPIServer(tt.fields.Storage, tt.fields.Worker)
			server := httptest.NewServer(api.Router)
			var url string
			if tt.args.IdOk != 0 {
				url = fmt.Sprintf("%s/api/fetcher/%v/history", server.URL, tt.args.IdOk)
			} else {
				url = fmt.Sprintf("%s/api/fetcher/%s/history", server.URL, tt.args.wrongId)
			}
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			resp := execReq(req, t)
			var respBody []HistoryElementResponse
			_ = json.NewDecoder(resp.Body).Decode(&respBody)
			if tt.args.wantCode != 0 {
				checkRespCode(t, tt.args.wantCode, resp.StatusCode)
			}
			if tt.args.wantCreatedAt != 0 {
				checkIntInBody(t, tt.args.wantCreatedAt, respBody[0].CreatedAt)
			}
		})
	}
}

func execReq(req *http.Request, t *testing.T) *http.Response {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf(err.Error())
	}
	return res
}
func checkRespCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected resp code %d. Got %d\n", expected, actual)
	}
}
func checkIntInBody(t *testing.T, expected int64, actual int64) {
	if expected != actual {
		t.Errorf("Expected int in resp body %d. Got %d\n", expected, actual)
	}
}
