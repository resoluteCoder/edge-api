// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/commits.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/redhatinsights/edge-api/pkg/models"
	reflect "reflect"
)

// MockCommitServiceInterface is a mock of CommitServiceInterface interface
type MockCommitServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCommitServiceInterfaceMockRecorder
}

// MockCommitServiceInterfaceMockRecorder is the mock recorder for MockCommitServiceInterface
type MockCommitServiceInterfaceMockRecorder struct {
	mock *MockCommitServiceInterface
}

// NewMockCommitServiceInterface creates a new mock instance
func NewMockCommitServiceInterface(ctrl *gomock.Controller) *MockCommitServiceInterface {
	mock := &MockCommitServiceInterface{ctrl: ctrl}
	mock.recorder = &MockCommitServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommitServiceInterface) EXPECT() *MockCommitServiceInterfaceMockRecorder {
	return m.recorder
}

// GetCommitByID mocks base method
func (m *MockCommitServiceInterface) GetCommitByID(commitID uint) (*models.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitByID", commitID)
	ret0, _ := ret[0].(*models.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitByID indicates an expected call of GetCommitByID
func (mr *MockCommitServiceInterfaceMockRecorder) GetCommitByID(commitID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitByID", reflect.TypeOf((*MockCommitServiceInterface)(nil).GetCommitByID), commitID)
}

// GetCommitByOSTreeCommit mocks base method
func (m *MockCommitServiceInterface) GetCommitByOSTreeCommit(ost string) (*models.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitByOSTreeCommit", ost)
	ret0, _ := ret[0].(*models.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitByOSTreeCommit indicates an expected call of GetCommitByOSTreeCommit
func (mr *MockCommitServiceInterfaceMockRecorder) GetCommitByOSTreeCommit(ost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitByOSTreeCommit", reflect.TypeOf((*MockCommitServiceInterface)(nil).GetCommitByOSTreeCommit), ost)
}
