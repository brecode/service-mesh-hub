// Code generated by MockGen. DO NOT EDIT.
// Source: ./validation.go

// Package mock_failover_service_validation is a generated GoMock package.
package mock_failover_service_validation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	failover "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/failover"
)

// MockFailoverServiceValidator is a mock of FailoverServiceValidator interface.
type MockFailoverServiceValidator struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverServiceValidatorMockRecorder
}

// MockFailoverServiceValidatorMockRecorder is the mock recorder for MockFailoverServiceValidator.
type MockFailoverServiceValidatorMockRecorder struct {
	mock *MockFailoverServiceValidator
}

// NewMockFailoverServiceValidator creates a new mock instance.
func NewMockFailoverServiceValidator(ctrl *gomock.Controller) *MockFailoverServiceValidator {
	mock := &MockFailoverServiceValidator{ctrl: ctrl}
	mock.recorder = &MockFailoverServiceValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFailoverServiceValidator) EXPECT() *MockFailoverServiceValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockFailoverServiceValidator) Validate(snapshot failover.InputSnapshot) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Validate", snapshot)
}

// Validate indicates an expected call of Validate.
func (mr *MockFailoverServiceValidatorMockRecorder) Validate(snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockFailoverServiceValidator)(nil).Validate), snapshot)
}