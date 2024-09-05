// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mock/service.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/nkuros/ebanxchallenge/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockAccountService) AddAccount(originId string, Balance int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAccount", originId, Balance)
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockAccountServiceMockRecorder) AddAccount(originId, Balance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockAccountService)(nil).AddAccount), originId, Balance)
}

// GetAccount mocks base method.
func (m *MockAccountService) GetAccount(originId string) (*entity.Account, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", originId)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountServiceMockRecorder) GetAccount(originId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountService)(nil).GetAccount), originId)
}
