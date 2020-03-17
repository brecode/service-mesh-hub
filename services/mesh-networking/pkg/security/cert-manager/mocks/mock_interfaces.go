// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_cert_manager is a generated GoMock package.
package mock_cert_manager

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1"
	types "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	types0 "github.com/solo-io/mesh-projects/pkg/api/security.zephyr.solo.io/v1alpha1/types"
)

// MockCertConfigProducer is a mock of CertConfigProducer interface
type MockCertConfigProducer struct {
	ctrl     *gomock.Controller
	recorder *MockCertConfigProducerMockRecorder
}

// MockCertConfigProducerMockRecorder is the mock recorder for MockCertConfigProducer
type MockCertConfigProducerMockRecorder struct {
	mock *MockCertConfigProducer
}

// NewMockCertConfigProducer creates a new mock instance
func NewMockCertConfigProducer(ctrl *gomock.Controller) *MockCertConfigProducer {
	mock := &MockCertConfigProducer{ctrl: ctrl}
	mock.recorder = &MockCertConfigProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertConfigProducer) EXPECT() *MockCertConfigProducerMockRecorder {
	return m.recorder
}

// ConfigureCertificateInfo mocks base method
func (m *MockCertConfigProducer) ConfigureCertificateInfo(mg *v1alpha10.MeshGroup, mesh *v1alpha1.Mesh) (*types0.CertConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureCertificateInfo", mg, mesh)
	ret0, _ := ret[0].(*types0.CertConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureCertificateInfo indicates an expected call of ConfigureCertificateInfo
func (mr *MockCertConfigProducerMockRecorder) ConfigureCertificateInfo(mg, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureCertificateInfo", reflect.TypeOf((*MockCertConfigProducer)(nil).ConfigureCertificateInfo), mg, mesh)
}

// MockMeshGroupCertificateManager is a mock of MeshGroupCertificateManager interface
type MockMeshGroupCertificateManager struct {
	ctrl     *gomock.Controller
	recorder *MockMeshGroupCertificateManagerMockRecorder
}

// MockMeshGroupCertificateManagerMockRecorder is the mock recorder for MockMeshGroupCertificateManager
type MockMeshGroupCertificateManagerMockRecorder struct {
	mock *MockMeshGroupCertificateManager
}

// NewMockMeshGroupCertificateManager creates a new mock instance
func NewMockMeshGroupCertificateManager(ctrl *gomock.Controller) *MockMeshGroupCertificateManager {
	mock := &MockMeshGroupCertificateManager{ctrl: ctrl}
	mock.recorder = &MockMeshGroupCertificateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshGroupCertificateManager) EXPECT() *MockMeshGroupCertificateManagerMockRecorder {
	return m.recorder
}

// InitializeCertificateForMeshGroup mocks base method
func (m *MockMeshGroupCertificateManager) InitializeCertificateForMeshGroup(ctx context.Context, new *v1alpha10.MeshGroup) types.MeshGroupStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeCertificateForMeshGroup", ctx, new)
	ret0, _ := ret[0].(types.MeshGroupStatus)
	return ret0
}

// InitializeCertificateForMeshGroup indicates an expected call of InitializeCertificateForMeshGroup
func (mr *MockMeshGroupCertificateManagerMockRecorder) InitializeCertificateForMeshGroup(ctx, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeCertificateForMeshGroup", reflect.TypeOf((*MockMeshGroupCertificateManager)(nil).InitializeCertificateForMeshGroup), ctx, new)
}