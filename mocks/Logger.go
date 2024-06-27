// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	constants "github.com/mmcloud/solder/pkg/constants"

	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: msg
func (_m *Logger) Debug(msg string) {
	_m.Called(msg)
}

// Debugf provides a mock function with given fields: format, args
func (_m *Logger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg
func (_m *Logger) Error(msg string) {
	_m.Called(msg)
}

// Errorf provides a mock function with given fields: format, args
func (_m *Logger) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: msg
func (_m *Logger) Fatal(msg string) {
	_m.Called(msg)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *Logger) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// GetLevel provides a mock function with given fields:
func (_m *Logger) GetLevel() constants.Level {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLevel")
	}

	var r0 constants.Level
	if rf, ok := ret.Get(0).(func() constants.Level); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(constants.Level)
	}

	return r0
}

// GetLoggerKind provides a mock function with given fields:
func (_m *Logger) GetLoggerKind() constants.LoggerKind {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLoggerKind")
	}

	var r0 constants.LoggerKind
	if rf, ok := ret.Get(0).(func() constants.LoggerKind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(constants.LoggerKind)
	}

	return r0
}

// Info provides a mock function with given fields: msg
func (_m *Logger) Info(msg string) {
	_m.Called(msg)
}

// Infof provides a mock function with given fields: format, args
func (_m *Logger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// SetLevel provides a mock function with given fields: level
func (_m *Logger) SetLevel(level constants.Level) {
	_m.Called(level)
}

// Warn provides a mock function with given fields: msg
func (_m *Logger) Warn(msg string) {
	_m.Called(msg)
}

// Warnf provides a mock function with given fields: format, args
func (_m *Logger) Warnf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}