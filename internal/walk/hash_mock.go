// Code generated by MockGen. DO NOT EDIT.
// Source: hash (interfaces: Hash)

// Package walk is a generated GoMock package.
package walk

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHash is a mock of Hash interface.
type MockHash struct {
	ctrl     *gomock.Controller
	recorder *MockHashMockRecorder
}

// MockHashMockRecorder is the mock recorder for MockHash.
type MockHashMockRecorder struct {
	mock *MockHash
}

// NewMockHash creates a new mock instance.
func NewMockHash(ctrl *gomock.Controller) *MockHash {
	mock := &MockHash{ctrl: ctrl}
	mock.recorder = &MockHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHash) EXPECT() *MockHashMockRecorder {
	return m.recorder
}

// BlockSize mocks base method.
func (m *MockHash) BlockSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// BlockSize indicates an expected call of BlockSize.
func (mr *MockHashMockRecorder) BlockSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSize", reflect.TypeOf((*MockHash)(nil).BlockSize))
}

// Reset mocks base method.
func (m *MockHash) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockHashMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockHash)(nil).Reset))
}

// Size mocks base method.
func (m *MockHash) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockHashMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockHash)(nil).Size))
}

// Sum mocks base method.
func (m *MockHash) Sum(arg0 []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Sum indicates an expected call of Sum.
func (mr *MockHashMockRecorder) Sum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockHash)(nil).Sum), arg0)
}

// Write mocks base method.
func (m *MockHash) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockHashMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockHash)(nil).Write), arg0)
}
