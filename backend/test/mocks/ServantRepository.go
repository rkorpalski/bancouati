// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import servant "github.com/codenation-dev/squad-4-aceleradev-fs-online-1/backend/pkg/servant"

// ServantRepository is an autogenerated mock type for the ServantRepository type
type ServantRepository struct {
	mock.Mock
}

// FindServantBySalary provides a mock function with given fields: salary
func (_m *ServantRepository) FindServantBySalary(salary float64) ([]servant.Servant, error) {
	ret := _m.Called(salary)

	var r0 []servant.Servant
	if rf, ok := ret.Get(0).(func(float64) []servant.Servant); ok {
		r0 = rf(salary)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]servant.Servant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float64) error); ok {
		r1 = rf(salary)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertClient provides a mock function with given fields: client
func (_m *ServantRepository) InsertClient(client servant.Client) error {
	ret := _m.Called(client)

	var r0 error
	if rf, ok := ret.Get(0).(func(servant.Client) error); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertServant provides a mock function with given fields: _a0
func (_m *ServantRepository) InsertServant(_a0 servant.Servant) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(servant.Servant) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsServantExists provides a mock function with given fields: _a0
func (_m *ServantRepository) IsServantExists(_a0 servant.Servant) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(servant.Servant) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(servant.Servant) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClient provides a mock function with given fields: client
func (_m *ServantRepository) UpdateClient(client string) error {
	ret := _m.Called(client)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSendAlert provides a mock function with given fields: _a0
func (_m *ServantRepository) UpdateSendAlert(_a0 servant.Servant) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(servant.Servant) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateServant provides a mock function with given fields: _a0
func (_m *ServantRepository) UpdateServant(_a0 servant.Servant) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(servant.Servant) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getPotentialClients provides a mock function with given fields:
func (_m *ServantRepository) getPotentialClients() ([]servant.Client, error) {
	ret := _m.Called()

	var r0 []servant.Client
	if rf, ok := ret.Get(0).(func() []servant.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]servant.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
