// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ChatData is an autogenerated mock type for the ChatData type
type ChatData struct {
	mock.Mock
}

type mockConstructorTestingTNewChatData interface {
	mock.TestingT
	Cleanup(func())
}

// NewChatData creates a new instance of ChatData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChatData(t mockConstructorTestingTNewChatData) *ChatData {
	mock := &ChatData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}