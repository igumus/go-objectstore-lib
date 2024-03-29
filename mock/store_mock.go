// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	objectstore "github.com/igumus/go-objectstore-lib"
	cid "github.com/ipfs/go-cid"
)

// MockObjectStore is a mock of ObjectStore interface.
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore.
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance.
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// CreateObject mocks base method.
func (m *MockObjectStore) CreateObject(arg0 context.Context, arg1 io.Reader) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObject", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObject indicates an expected call of CreateObject.
func (mr *MockObjectStoreMockRecorder) CreateObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObject", reflect.TypeOf((*MockObjectStore)(nil).CreateObject), arg0, arg1)
}

// HasObject mocks base method.
func (m *MockObjectStore) HasObject(arg0 context.Context, arg1 cid.Cid) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasObject", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasObject indicates an expected call of HasObject.
func (mr *MockObjectStoreMockRecorder) HasObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasObject", reflect.TypeOf((*MockObjectStore)(nil).HasObject), arg0, arg1)
}

// ListObject mocks base method.
func (m *MockObjectStore) ListObject(arg0 context.Context) <-chan objectstore.ListObjectEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObject", arg0)
	ret0, _ := ret[0].(<-chan objectstore.ListObjectEvent)
	return ret0
}

// ListObject indicates an expected call of ListObject.
func (mr *MockObjectStoreMockRecorder) ListObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObject", reflect.TypeOf((*MockObjectStore)(nil).ListObject), arg0)
}

// ReadObject mocks base method.
func (m *MockObjectStore) ReadObject(arg0 context.Context, arg1 cid.Cid) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadObject", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadObject indicates an expected call of ReadObject.
func (mr *MockObjectStoreMockRecorder) ReadObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadObject", reflect.TypeOf((*MockObjectStore)(nil).ReadObject), arg0, arg1)
}
