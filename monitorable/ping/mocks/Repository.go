// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import (
	models "github.com/monitoror/monitoror/monitorable/ping/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ExecutePing provides a mock function with given fields: hostname
func (_m *Repository) ExecutePing(hostname string) (*models.Ping, error) {
	ret := _m.Called(hostname)

	var r0 *models.Ping
	if rf, ok := ret.Get(0).(func(string) *models.Ping); ok {
		r0 = rf(hostname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Ping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hostname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
