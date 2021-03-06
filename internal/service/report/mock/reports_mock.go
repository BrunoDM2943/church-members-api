// Code generated by MockGen. DO NOT EDIT.
// Source: ./reportsService.go

// Package mock_report is a generated GoMock package.
package mock_report

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// BirthdayReport mocks base method.
func (m *MockService) BirthdayReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BirthdayReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BirthdayReport indicates an expected call of BirthdayReport.
func (mr *MockServiceMockRecorder) BirthdayReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BirthdayReport", reflect.TypeOf((*MockService)(nil).BirthdayReport))
}

// ClassificationReport mocks base method.
func (m *MockService) ClassificationReport(classification string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassificationReport", classification)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassificationReport indicates an expected call of ClassificationReport.
func (mr *MockServiceMockRecorder) ClassificationReport(classification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassificationReport", reflect.TypeOf((*MockService)(nil).ClassificationReport), classification)
}

// LegalReport mocks base method.
func (m *MockService) LegalReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegalReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LegalReport indicates an expected call of LegalReport.
func (mr *MockServiceMockRecorder) LegalReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegalReport", reflect.TypeOf((*MockService)(nil).LegalReport))
}

// MarriageReport mocks base method.
func (m *MockService) MarriageReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarriageReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarriageReport indicates an expected call of MarriageReport.
func (mr *MockServiceMockRecorder) MarriageReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarriageReport", reflect.TypeOf((*MockService)(nil).MarriageReport))
}

// MemberReport mocks base method.
func (m *MockService) MemberReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberReport indicates an expected call of MemberReport.
func (mr *MockServiceMockRecorder) MemberReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberReport", reflect.TypeOf((*MockService)(nil).MemberReport))
}
