// Code generated by MockGen. DO NOT EDIT.
// Source: ./reports.go

// Package mock_reports is a generated GoMock package.
package mock_reports

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReportsGenerator is a mock of ReportsGenerator interface
type MockReportsGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockReportsGeneratorMockRecorder
}

// MockReportsGeneratorMockRecorder is the mock recorder for MockReportsGenerator
type MockReportsGeneratorMockRecorder struct {
	mock *MockReportsGenerator
}

// NewMockReportsGenerator creates a new mock instance
func NewMockReportsGenerator(ctrl *gomock.Controller) *MockReportsGenerator {
	mock := &MockReportsGenerator{ctrl: ctrl}
	mock.recorder = &MockReportsGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReportsGenerator) EXPECT() *MockReportsGeneratorMockRecorder {
	return m.recorder
}

// LegalReport mocks base method
func (m *MockReportsGenerator) LegalReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegalReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LegalReport indicates an expected call of LegalReport
func (mr *MockReportsGeneratorMockRecorder) LegalReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegalReport", reflect.TypeOf((*MockReportsGenerator)(nil).LegalReport))
}

// MemberReport mocks base method
func (m *MockReportsGenerator) MemberReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MemberReport indicates an expected call of MemberReport
func (mr *MockReportsGeneratorMockRecorder) MemberReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberReport", reflect.TypeOf((*MockReportsGenerator)(nil).MemberReport))
}

// BirthdayReport mocks base method
func (m *MockReportsGenerator) BirthdayReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BirthdayReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BirthdayReport indicates an expected call of BirthdayReport
func (mr *MockReportsGeneratorMockRecorder) BirthdayReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BirthdayReport", reflect.TypeOf((*MockReportsGenerator)(nil).BirthdayReport))
}

// MariageReport mocks base method
func (m *MockReportsGenerator) MariageReport() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MariageReport")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MariageReport indicates an expected call of MariageReport
func (mr *MockReportsGeneratorMockRecorder) MariageReport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MariageReport", reflect.TypeOf((*MockReportsGenerator)(nil).MariageReport))
}