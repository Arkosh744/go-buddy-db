// Code generated by MockGen. DO NOT EDIT.
// Source: compute.go

// Package compute is a generated GoMock package.
package compute

import (
	reflect "reflect"

	models "github.com/Arkosh744/go-buddy-db/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// Mockparser is a mock of parser interface.
type Mockparser struct {
	ctrl     *gomock.Controller
	recorder *MockparserMockRecorder
}

// MockparserMockRecorder is the mock recorder for Mockparser.
type MockparserMockRecorder struct {
	mock *Mockparser
}

// NewMockparser creates a new mock instance.
func NewMockparser(ctrl *gomock.Controller) *Mockparser {
	mock := &Mockparser{ctrl: ctrl}
	mock.recorder = &MockparserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockparser) EXPECT() *MockparserMockRecorder {
	return m.recorder
}

// ParseQuery mocks base method.
func (m *Mockparser) ParseQuery(query string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseQuery", query)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseQuery indicates an expected call of ParseQuery.
func (mr *MockparserMockRecorder) ParseQuery(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseQuery", reflect.TypeOf((*Mockparser)(nil).ParseQuery), query)
}

// Mockanalyzer is a mock of analyzer interface.
type Mockanalyzer struct {
	ctrl     *gomock.Controller
	recorder *MockanalyzerMockRecorder
}

// MockanalyzerMockRecorder is the mock recorder for Mockanalyzer.
type MockanalyzerMockRecorder struct {
	mock *Mockanalyzer
}

// NewMockanalyzer creates a new mock instance.
func NewMockanalyzer(ctrl *gomock.Controller) *Mockanalyzer {
	mock := &Mockanalyzer{ctrl: ctrl}
	mock.recorder = &MockanalyzerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockanalyzer) EXPECT() *MockanalyzerMockRecorder {
	return m.recorder
}

// AnalyzeQuery mocks base method.
func (m *Mockanalyzer) AnalyzeQuery(tokens []string) (models.Query, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnalyzeQuery", tokens)
	ret0, _ := ret[0].(models.Query)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnalyzeQuery indicates an expected call of AnalyzeQuery.
func (mr *MockanalyzerMockRecorder) AnalyzeQuery(tokens interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnalyzeQuery", reflect.TypeOf((*Mockanalyzer)(nil).AnalyzeQuery), tokens)
}
