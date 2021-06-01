// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-progress-api/core/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	progress "github.com/ozoncp/ocp-progress-api/core/progress"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddProgress mocks base method.
func (m *MockRepo) AddProgress(arg0 []progress.Pogress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProgress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProgress indicates an expected call of AddProgress.
func (mr *MockRepoMockRecorder) AddProgress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProgress", reflect.TypeOf((*MockRepo)(nil).AddProgress), arg0)
}

// DescribeProgress mocks base method.
func (m *MockRepo) DescribeProgress(arg0 uint64) (*progress.Pogress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProgress", arg0)
	ret0, _ := ret[0].(*progress.Pogress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProgress indicates an expected call of DescribeProgress.
func (mr *MockRepoMockRecorder) DescribeProgress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProgress", reflect.TypeOf((*MockRepo)(nil).DescribeProgress), arg0)
}

// ListProgress mocks base method.
func (m *MockRepo) ListProgress(arg0, arg1 uint64) ([]progress.Pogress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProgress", arg0, arg1)
	ret0, _ := ret[0].([]progress.Pogress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProgress indicates an expected call of ListProgress.
func (mr *MockRepoMockRecorder) ListProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProgress", reflect.TypeOf((*MockRepo)(nil).ListProgress), arg0, arg1)
}

// RemoveProgress mocks base method.
func (m *MockRepo) RemoveProgress(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProgress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProgress indicates an expected call of RemoveProgress.
func (mr *MockRepoMockRecorder) RemoveProgress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProgress", reflect.TypeOf((*MockRepo)(nil).RemoveProgress), arg0)
}
