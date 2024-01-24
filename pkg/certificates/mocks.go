// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/oidc-apps-controller/pkg/certificates (interfaces: CertificateOperations)

// Package certificates is a generated GoMock package.
package certificates

import (
	rsa "crypto/rsa"
	x509 "crypto/x509"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCertificateOperations is a mock of CertificateOperations interface.
type MockCertificateOperations struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateOperationsMockRecorder
}

// MockCertificateOperationsMockRecorder is the mock recorder for MockCertificateOperations.
type MockCertificateOperationsMockRecorder struct {
	mock *MockCertificateOperations
}

// NewMockCertificateOperations creates a new mock instance.
func NewMockCertificateOperations(ctrl *gomock.Controller) *MockCertificateOperations {
	mock := &MockCertificateOperations{ctrl: ctrl}
	mock.recorder = &MockCertificateOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateOperations) EXPECT() *MockCertificateOperationsMockRecorder {
	return m.recorder
}

// CreateCertificate mocks base method.
func (m *MockCertificateOperations) CreateCertificate(arg0 io.Reader, arg1, arg2 *x509.Certificate, arg3, arg4 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCertificate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCertificate indicates an expected call of CreateCertificate.
func (mr *MockCertificateOperationsMockRecorder) CreateCertificate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificate", reflect.TypeOf((*MockCertificateOperations)(nil).CreateCertificate), arg0, arg1, arg2, arg3, arg4)
}

// GenerateKey mocks base method.
func (m *MockCertificateOperations) GenerateKey(arg0 io.Reader, arg1 int) (*rsa.PrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKey", arg0, arg1)
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKey indicates an expected call of GenerateKey.
func (mr *MockCertificateOperationsMockRecorder) GenerateKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKey", reflect.TypeOf((*MockCertificateOperations)(nil).GenerateKey), arg0, arg1)
}