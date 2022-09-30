// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AssignApplicationToFormation provides a mock function with given fields: appId, formationName
func (_m *Client) AssignApplicationToFormation(appId string, formationName string) error {
	ret := _m.Called(appId, formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(appId, formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssignRuntimeToFormation provides a mock function with given fields: runtimeId, formationName
func (_m *Client) AssignRuntimeToFormation(runtimeId string, formationName string) error {
	ret := _m.Called(runtimeId, formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(runtimeId, formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConnectionToken provides a mock function with given fields: runtimeID
func (_m *Client) GetConnectionToken(runtimeID string) (string, string, error) {
	ret := _m.Called(runtimeID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(runtimeID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(runtimeID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterApplication provides a mock function with given fields: appName, scenario
func (_m *Client) RegisterApplication(appName string, scenario string) (string, error) {
	ret := _m.Called(appName, scenario)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(appName, scenario)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(appName, scenario)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterFormation provides a mock function with given fields: formationName
func (_m *Client) RegisterFormation(formationName string) error {
	ret := _m.Called(formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterRuntime provides a mock function with given fields: runtimeName
func (_m *Client) RegisterRuntime(runtimeName string) (string, error) {
	ret := _m.Called(runtimeName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(runtimeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(runtimeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterApplication provides a mock function with given fields: id
func (_m *Client) UnregisterApplication(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnregisterFormation provides a mock function with given fields: formationName
func (_m *Client) UnregisterFormation(formationName string) error {
	ret := _m.Called(formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnregisterRuntime provides a mock function with given fields: id
func (_m *Client) UnregisterRuntime(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}