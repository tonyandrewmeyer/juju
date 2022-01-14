// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: ModelOperation)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	txn "github.com/juju/mgo/v2/txn"
)

// MockModelOperation is a mock of ModelOperation interface.
type MockModelOperation struct {
	ctrl     *gomock.Controller
	recorder *MockModelOperationMockRecorder
}

// MockModelOperationMockRecorder is the mock recorder for MockModelOperation.
type MockModelOperationMockRecorder struct {
	mock *MockModelOperation
}

// NewMockModelOperation creates a new mock instance.
func NewMockModelOperation(ctrl *gomock.Controller) *MockModelOperation {
	mock := &MockModelOperation{ctrl: ctrl}
	mock.recorder = &MockModelOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelOperation) EXPECT() *MockModelOperationMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockModelOperation) Build(arg0 int) ([]txn.Op, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].([]txn.Op)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockModelOperationMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockModelOperation)(nil).Build), arg0)
}

// Done mocks base method.
func (m *MockModelOperation) Done(arg0 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockModelOperationMockRecorder) Done(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockModelOperation)(nil).Done), arg0)
}
