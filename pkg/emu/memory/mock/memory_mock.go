// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/emu/memory/memory.go

// Package mock_memory is a generated GoMock package.
package mock_memory

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReaderWriter is a mock of ReaderWriter interface
type MockReaderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockReaderWriterMockRecorder
}

// MockReaderWriterMockRecorder is the mock recorder for MockReaderWriter
type MockReaderWriterMockRecorder struct {
	mock *MockReaderWriter
}

// NewMockReaderWriter creates a new mock instance
func NewMockReaderWriter(ctrl *gomock.Controller) *MockReaderWriter {
	mock := &MockReaderWriter{ctrl: ctrl}
	mock.recorder = &MockReaderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReaderWriter) EXPECT() *MockReaderWriterMockRecorder {
	return m.recorder
}

// ReadValue mocks base method
func (m *MockReaderWriter) ReadValue(addr uint16) uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadValue", addr)
	ret0, _ := ret[0].(uint8)
	return ret0
}

// ReadValue indicates an expected call of ReadValue
func (mr *MockReaderWriterMockRecorder) ReadValue(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadValue", reflect.TypeOf((*MockReaderWriter)(nil).ReadValue), addr)
}

// WriteValue mocks base method
func (m *MockReaderWriter) WriteValue(addr uint16, value uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteValue", addr, value)
}

// WriteValue indicates an expected call of WriteValue
func (mr *MockReaderWriterMockRecorder) WriteValue(addr, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteValue", reflect.TypeOf((*MockReaderWriter)(nil).WriteValue), addr, value)
}

// ReadBytes mocks base method
func (m *MockReaderWriter) ReadBytes(start uint16, n uint8) []uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBytes", start, n)
	ret0, _ := ret[0].([]uint8)
	return ret0
}

// ReadBytes indicates an expected call of ReadBytes
func (mr *MockReaderWriterMockRecorder) ReadBytes(start, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBytes", reflect.TypeOf((*MockReaderWriter)(nil).ReadBytes), start, n)
}

// WriteBytes mocks base method
func (m *MockReaderWriter) WriteBytes(start uint16, data []uint8) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteBytes", start, data)
}

// WriteBytes indicates an expected call of WriteBytes
func (mr *MockReaderWriterMockRecorder) WriteBytes(start, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBytes", reflect.TypeOf((*MockReaderWriter)(nil).WriteBytes), start, data)
}

// GetHexDigitAddress mocks base method
func (m *MockReaderWriter) GetHexDigitAddress(digit uint8) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHexDigitAddress", digit)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// GetHexDigitAddress indicates an expected call of GetHexDigitAddress
func (mr *MockReaderWriterMockRecorder) GetHexDigitAddress(digit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHexDigitAddress", reflect.TypeOf((*MockReaderWriter)(nil).GetHexDigitAddress), digit)
}
