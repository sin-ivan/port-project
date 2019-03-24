// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	port "github.com/port-project/domain-service/port"
	reflect "reflect"
)

// MockPortRepository is a mock of PortRepository interface
type MockPortRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPortRepositoryMockRecorder
}

// MockPortRepositoryMockRecorder is the mock recorder for MockPortRepository
type MockPortRepositoryMockRecorder struct {
	mock *MockPortRepository
}

// NewMockPortRepository creates a new mock instance
func NewMockPortRepository(ctrl *gomock.Controller) *MockPortRepository {
	mock := &MockPortRepository{ctrl: ctrl}
	mock.recorder = &MockPortRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortRepository) EXPECT() *MockPortRepositoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockPortRepository) GetByID(arg0 string) (*port.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*port.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPortRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPortRepository)(nil).GetByID), arg0)
}

// GetAll mocks base method
func (m *MockPortRepository) GetAll() ([]*port.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*port.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPortRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPortRepository)(nil).GetAll))
}

// Store mocks base method
func (m *MockPortRepository) Store(p *port.Port) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockPortRepositoryMockRecorder) Store(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockPortRepository)(nil).Store), p)
}

// Update mocks base method
func (m *MockPortRepository) Update(p *port.Port) (*port.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", p)
	ret0, _ := ret[0].(*port.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockPortRepositoryMockRecorder) Update(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPortRepository)(nil).Update), p)
}

// Delete mocks base method
func (m *MockPortRepository) Delete(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockPortRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPortRepository)(nil).Delete), id)
}
