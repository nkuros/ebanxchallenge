// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -source=controller.go -destination=mock/controller.go
//

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAccountController is a mock of AccountController interface.
type MockAccountController struct {
	ctrl     *gomock.Controller
	recorder *MockAccountControllerMockRecorder
}

// MockAccountControllerMockRecorder is the mock recorder for MockAccountController.
type MockAccountControllerMockRecorder struct {
	mock *MockAccountController
}

// NewMockAccountController creates a new mock instance.
func NewMockAccountController(ctrl *gomock.Controller) *MockAccountController {
	mock := &MockAccountController{ctrl: ctrl}
	mock.recorder = &MockAccountControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountController) EXPECT() *MockAccountControllerMockRecorder {
	return m.recorder
}

// DeleteAllAccountsController mocks base method.
func (m *MockAccountController) DeleteAllAccountsController() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteAllAccountsController")
}

// DeleteAllAccountsController indicates an expected call of DeleteAllAccountsController.
func (mr *MockAccountControllerMockRecorder) DeleteAllAccountsController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllAccountsController", reflect.TypeOf((*MockAccountController)(nil).DeleteAllAccountsController))
}

// GetBalanceController mocks base method.
func (m *MockAccountController) GetBalanceController(originId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceController", originId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceController indicates an expected call of GetBalanceController.
func (mr *MockAccountControllerMockRecorder) GetBalanceController(originId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceController", reflect.TypeOf((*MockAccountController)(nil).GetBalanceController), originId)
}

// PostDepositEventController mocks base method.
func (m *MockAccountController) PostDepositEventController(originId string, amount int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDepositEventController", originId, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostDepositEventController indicates an expected call of PostDepositEventController.
func (mr *MockAccountControllerMockRecorder) PostDepositEventController(originId, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDepositEventController", reflect.TypeOf((*MockAccountController)(nil).PostDepositEventController), originId, amount)
}

// PostTransferEventController mocks base method.
func (m *MockAccountController) PostTransferEventController(originId, targetId string, amount int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostTransferEventController", originId, targetId, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostTransferEventController indicates an expected call of PostTransferEventController.
func (mr *MockAccountControllerMockRecorder) PostTransferEventController(originId, targetId, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostTransferEventController", reflect.TypeOf((*MockAccountController)(nil).PostTransferEventController), originId, targetId, amount)
}

// PostWithdrawEventController mocks base method.
func (m *MockAccountController) PostWithdrawEventController(originId string, amount int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWithdrawEventController", originId, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostWithdrawEventController indicates an expected call of PostWithdrawEventController.
func (mr *MockAccountControllerMockRecorder) PostWithdrawEventController(originId, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWithdrawEventController", reflect.TypeOf((*MockAccountController)(nil).PostWithdrawEventController), originId, amount)
}
