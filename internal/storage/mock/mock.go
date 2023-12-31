// Code generated by MockGen. DO NOT EDIT.
// Source: storage_gofermat.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	handlersmodels "gofermart/internal/models/handlers_models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorageGofermart is a mock of StorageGofermart interface.
type MockStorageGofermart struct {
	ctrl     *gomock.Controller
	recorder *MockStorageGofermartMockRecorder
}

// MockStorageGofermartMockRecorder is the mock recorder for MockStorageGofermart.
type MockStorageGofermartMockRecorder struct {
	mock *MockStorageGofermart
}

// NewMockStorageGofermart creates a new mock instance.
func NewMockStorageGofermart(ctrl *gomock.Controller) *MockStorageGofermart {
	mock := &MockStorageGofermart{ctrl: ctrl}
	mock.recorder = &MockStorageGofermartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageGofermart) EXPECT() *MockStorageGofermartMockRecorder {
	return m.recorder
}

// AddNewOrderAndAccrual mocks base method.
func (m *MockStorageGofermart) AddNewOrderAndAccrual(ctxRequest context.Context, reqOrder *handlersmodels.ReqOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewOrderAndAccrual", ctxRequest, reqOrder)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewOrderAndAccrual indicates an expected call of AddNewOrderAndAccrual.
func (mr *MockStorageGofermartMockRecorder) AddNewOrderAndAccrual(ctxRequest, reqOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewOrderAndAccrual", reflect.TypeOf((*MockStorageGofermart)(nil).AddNewOrderAndAccrual), ctxRequest, reqOrder)
}

// AddNewUserAndBalance mocks base method.
func (m *MockStorageGofermart) AddNewUserAndBalance(ctxRequest context.Context, reqRegister handlersmodels.RequestRegister) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUserAndBalance", ctxRequest, reqRegister)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewUserAndBalance indicates an expected call of AddNewUserAndBalance.
func (mr *MockStorageGofermartMockRecorder) AddNewUserAndBalance(ctxRequest, reqRegister interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUserAndBalance", reflect.TypeOf((*MockStorageGofermart)(nil).AddNewUserAndBalance), ctxRequest, reqRegister)
}

// CheckUserLoginData mocks base method.
func (m *MockStorageGofermart) CheckUserLoginData(reqLogin *handlersmodels.RequestLogin) (*handlersmodels.ResultLogin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserLoginData", reqLogin)
	ret0, _ := ret[0].(*handlersmodels.ResultLogin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserLoginData indicates an expected call of CheckUserLoginData.
func (mr *MockStorageGofermartMockRecorder) CheckUserLoginData(reqLogin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserLoginData", reflect.TypeOf((*MockStorageGofermart)(nil).CheckUserLoginData), reqLogin)
}

// CloseDB mocks base method.
func (m *MockStorageGofermart) CloseDB() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseDB")
}

// CloseDB indicates an expected call of CloseDB.
func (mr *MockStorageGofermartMockRecorder) CloseDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDB", reflect.TypeOf((*MockStorageGofermart)(nil).CloseDB))
}

// GetAllHistoryBalance mocks base method.
func (m *MockStorageGofermart) GetAllHistoryBalance(ctx context.Context, userID int) ([]handlersmodels.RespWithdrawalsHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllHistoryBalance", ctx, userID)
	ret0, _ := ret[0].([]handlersmodels.RespWithdrawalsHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllHistoryBalance indicates an expected call of GetAllHistoryBalance.
func (mr *MockStorageGofermartMockRecorder) GetAllHistoryBalance(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllHistoryBalance", reflect.TypeOf((*MockStorageGofermart)(nil).GetAllHistoryBalance), ctx, userID)
}

// GetAllUserOrders mocks base method.
func (m *MockStorageGofermart) GetAllUserOrders(ctx context.Context, userID int) ([]handlersmodels.RespGetOrders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserOrders", ctx, userID)
	ret0, _ := ret[0].([]handlersmodels.RespGetOrders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserOrders indicates an expected call of GetAllUserOrders.
func (mr *MockStorageGofermartMockRecorder) GetAllUserOrders(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserOrders", reflect.TypeOf((*MockStorageGofermart)(nil).GetAllUserOrders), ctx, userID)
}

// GetUserBalance mocks base method.
func (m *MockStorageGofermart) GetUserBalance(ctx context.Context) (*handlersmodels.RespUserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalance", ctx)
	ret0, _ := ret[0].(*handlersmodels.RespUserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalance indicates an expected call of GetUserBalance.
func (mr *MockStorageGofermartMockRecorder) GetUserBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalance", reflect.TypeOf((*MockStorageGofermart)(nil).GetUserBalance), ctx)
}

// InTransaction mocks base method.
func (m *MockStorageGofermart) InTransaction(parentsCtx context.Context, f func(context.Context, *sql.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InTransaction", parentsCtx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// InTransaction indicates an expected call of InTransaction.
func (mr *MockStorageGofermartMockRecorder) InTransaction(parentsCtx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InTransaction", reflect.TypeOf((*MockStorageGofermart)(nil).InTransaction), parentsCtx, f)
}

// ProcessingDebitingFunds mocks base method.
func (m *MockStorageGofermart) ProcessingDebitingFunds(ctxRequest context.Context, reqWithdraw handlersmodels.ReqWithdraw) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessingDebitingFunds", ctxRequest, reqWithdraw)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessingDebitingFunds indicates an expected call of ProcessingDebitingFunds.
func (mr *MockStorageGofermartMockRecorder) ProcessingDebitingFunds(ctxRequest, reqWithdraw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessingDebitingFunds", reflect.TypeOf((*MockStorageGofermart)(nil).ProcessingDebitingFunds), ctxRequest, reqWithdraw)
}
