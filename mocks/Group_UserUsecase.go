// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "lesgoobackend/domain"

	messaging "firebase.google.com/go/messaging"

	mock "github.com/stretchr/testify/mock"
)

// Group_UserUsecase is an autogenerated mock type for the Group_UserUsecase type
type Group_UserUsecase struct {
	mock.Mock
}

// AddJoined provides a mock function with given fields: data
func (_m *Group_UserUsecase) AddJoined(data domain.Group_User) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Group_User) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LeaveGroup provides a mock function with given fields: data
func (_m *Group_UserUsecase) LeaveGroup(data domain.Group_User) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Group_User) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLocation provides a mock function with given fields: data, client, _a2
func (_m *Group_UserUsecase) UpdateLocation(data domain.Group_User, client *messaging.Client, _a2 context.Context) error {
	ret := _m.Called(data, client, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Group_User, *messaging.Client, context.Context) error); ok {
		r0 = rf(data, client, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewGroup_UserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroup_UserUsecase creates a new instance of Group_UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroup_UserUsecase(t mockConstructorTestingTNewGroup_UserUsecase) *Group_UserUsecase {
	mock := &Group_UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
