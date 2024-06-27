// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/mmcloud/solder/pkg/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// GetLoggingConfig provides a mock function with given fields:
func (_m *Config) GetLoggingConfig() interfaces.LoggingConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLoggingConfig")
	}

	var r0 interfaces.LoggingConfig
	if rf, ok := ret.Get(0).(func() interfaces.LoggingConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.LoggingConfig)
		}
	}

	return r0
}

// GetServiceConfig provides a mock function with given fields:
func (_m *Config) GetServiceConfig() interfaces.ServiceConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetServiceConfig")
	}

	var r0 interfaces.ServiceConfig
	if rf, ok := ret.Get(0).(func() interfaces.ServiceConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.ServiceConfig)
		}
	}

	return r0
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}