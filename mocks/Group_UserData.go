// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Group_UserData is an autogenerated mock type for the Group_UserData type
type Group_UserData struct {
	mock.Mock
}

type mockConstructorTestingTNewGroup_UserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroup_UserData creates a new instance of Group_UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroup_UserData(t mockConstructorTestingTNewGroup_UserData) *Group_UserData {
	mock := &Group_UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
