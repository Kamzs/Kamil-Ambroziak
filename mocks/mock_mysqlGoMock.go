// Code generated by MockGen. DO NOT EDIT.
// Source: fetchers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	Kamil_Ambroziak "github.com/Kamzs/Kamil-Ambroziak"
	utils "github.com/Kamzs/Kamil-Ambroziak/utils"
	gomock "github.com/golang/mock/gomock"
	v3 "github.com/robfig/cron/v3"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// DeleteFetcher mocks base method.
func (m *MockStorage) DeleteFetcher(fetcherId int64) utils.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFetcher", fetcherId)
	ret0, _ := ret[0].(utils.RestErr)
	return ret0
}

// DeleteFetcher indicates an expected call of DeleteFetcher.
func (mr *MockStorageMockRecorder) DeleteFetcher(fetcherId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFetcher", reflect.TypeOf((*MockStorage)(nil).DeleteFetcher), fetcherId)
}

// FindAllFetchers mocks base method.
func (m *MockStorage) FindAllFetchers() ([]Kamil_Ambroziak.Fetcher, utils.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFetchers")
	ret0, _ := ret[0].([]Kamil_Ambroziak.Fetcher)
	ret1, _ := ret[1].(utils.RestErr)
	return ret0, ret1
}

// FindAllFetchers indicates an expected call of FindAllFetchers.
func (mr *MockStorageMockRecorder) FindAllFetchers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFetchers", reflect.TypeOf((*MockStorage)(nil).FindAllFetchers))
}

// GetFetcher mocks base method.
func (m *MockStorage) GetFetcher(fetcherId int64) (*Kamil_Ambroziak.Fetcher, utils.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFetcher", fetcherId)
	ret0, _ := ret[0].(*Kamil_Ambroziak.Fetcher)
	ret1, _ := ret[1].(utils.RestErr)
	return ret0, ret1
}

// GetFetcher indicates an expected call of GetFetcher.
func (mr *MockStorageMockRecorder) GetFetcher(fetcherId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFetcher", reflect.TypeOf((*MockStorage)(nil).GetFetcher), fetcherId)
}

// GetHistoryForFetcher mocks base method.
func (m *MockStorage) GetHistoryForFetcher(fetcherId int64) ([]Kamil_Ambroziak.HistoryElement, utils.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryForFetcher", fetcherId)
	ret0, _ := ret[0].([]Kamil_Ambroziak.HistoryElement)
	ret1, _ := ret[1].(utils.RestErr)
	return ret0, ret1
}

// GetHistoryForFetcher indicates an expected call of GetHistoryForFetcher.
func (mr *MockStorageMockRecorder) GetHistoryForFetcher(fetcherId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryForFetcher", reflect.TypeOf((*MockStorage)(nil).GetHistoryForFetcher), fetcherId)
}

// SaveFetcher mocks base method.
func (m *MockStorage) SaveFetcher(fetcher *Kamil_Ambroziak.Fetcher) utils.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFetcher", fetcher)
	ret0, _ := ret[0].(utils.RestErr)
	return ret0
}

// SaveFetcher indicates an expected call of SaveFetcher.
func (mr *MockStorageMockRecorder) SaveFetcher(fetcher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFetcher", reflect.TypeOf((*MockStorage)(nil).SaveFetcher), fetcher)
}

// SaveHistoryForFetcher mocks base method.
func (m *MockStorage) SaveHistoryForFetcher(historyEl *Kamil_Ambroziak.HistoryElement) utils.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveHistoryForFetcher", historyEl)
	ret0, _ := ret[0].(utils.RestErr)
	return ret0
}

// SaveHistoryForFetcher indicates an expected call of SaveHistoryForFetcher.
func (mr *MockStorageMockRecorder) SaveHistoryForFetcher(historyEl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveHistoryForFetcher", reflect.TypeOf((*MockStorage)(nil).SaveHistoryForFetcher), historyEl)
}

// UpdateFetcher mocks base method.
func (m *MockStorage) UpdateFetcher(fetcher *Kamil_Ambroziak.Fetcher) utils.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFetcher", fetcher)
	ret0, _ := ret[0].(utils.RestErr)
	return ret0
}

// UpdateFetcher indicates an expected call of UpdateFetcher.
func (mr *MockStorageMockRecorder) UpdateFetcher(fetcher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFetcher", reflect.TypeOf((*MockStorage)(nil).UpdateFetcher), fetcher)
}

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// DeregisterFetcher mocks base method.
func (m *MockWorker) DeregisterFetcher(jobId v3.EntryID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeregisterFetcher", jobId)
}

// DeregisterFetcher indicates an expected call of DeregisterFetcher.
func (mr *MockWorkerMockRecorder) DeregisterFetcher(jobId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterFetcher", reflect.TypeOf((*MockWorker)(nil).DeregisterFetcher), jobId)
}

// RegisterFetcher mocks base method.
func (m *MockWorker) RegisterFetcher(fetcher *Kamil_Ambroziak.Fetcher) (v3.EntryID, utils.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFetcher", fetcher)
	ret0, _ := ret[0].(v3.EntryID)
	ret1, _ := ret[1].(utils.RestErr)
	return ret0, ret1
}

// RegisterFetcher indicates an expected call of RegisterFetcher.
func (mr *MockWorkerMockRecorder) RegisterFetcher(fetcher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFetcher", reflect.TypeOf((*MockWorker)(nil).RegisterFetcher), fetcher)
}