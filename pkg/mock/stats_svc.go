// Code generated by MockGen. DO NOT EDIT.
// Source: stats.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStatsServicer is a mock of StatsServicer interface.
type MockStatsServicer struct {
	ctrl     *gomock.Controller
	recorder *MockStatsServicerMockRecorder
}

// MockStatsServicerMockRecorder is the mock recorder for MockStatsServicer.
type MockStatsServicerMockRecorder struct {
	mock *MockStatsServicer
}

// NewMockStatsServicer creates a new mock instance.
func NewMockStatsServicer(ctrl *gomock.Controller) *MockStatsServicer {
	mock := &MockStatsServicer{ctrl: ctrl}
	mock.recorder = &MockStatsServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsServicer) EXPECT() *MockStatsServicerMockRecorder {
	return m.recorder
}

// GetMaxEntries mocks base method.
func (m *MockStatsServicer) GetMaxEntries(ctx context.Context) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxEntries", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMaxEntries indicates an expected call of GetMaxEntries.
func (mr *MockStatsServicerMockRecorder) GetMaxEntries(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxEntries", reflect.TypeOf((*MockStatsServicer)(nil).GetMaxEntries), ctx)
}

// Health mocks base method.
func (m *MockStatsServicer) Health(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockStatsServicerMockRecorder) Health(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockStatsServicer)(nil).Health), ctx)
}

// IncrementCount mocks base method.
func (m *MockStatsServicer) IncrementCount(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementCount", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementCount indicates an expected call of IncrementCount.
func (mr *MockStatsServicerMockRecorder) IncrementCount(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementCount", reflect.TypeOf((*MockStatsServicer)(nil).IncrementCount), ctx, key)
}

// Reset mocks base method.
func (m *MockStatsServicer) Reset(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockStatsServicerMockRecorder) Reset(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockStatsServicer)(nil).Reset), ctx)
}
