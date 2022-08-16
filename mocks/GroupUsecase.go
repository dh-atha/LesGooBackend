// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "lesgoobackend/domain"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	session "github.com/aws/aws-sdk-go/aws/session"
)

// GroupUsecase is an autogenerated mock type for the GroupUsecase type
type GroupUsecase struct {
	mock.Mock
}

// AddGroup provides a mock function with given fields: data
func (_m *GroupUsecase) AddGroup(data domain.Group) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Group) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddGroupUser provides a mock function with given fields: dataUser
func (_m *GroupUsecase) AddGroupUser(dataUser domain.Group_User) error {
	ret := _m.Called(dataUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Group_User) error); ok {
		r0 = rf(dataUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroupByID provides a mock function with given fields: id, id_user
func (_m *GroupUsecase) DeleteGroupByID(id string, id_user uint) error {
	ret := _m.Called(id, id_user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(id, id_user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetChatsAndUsersLocation provides a mock function with given fields: groupID
func (_m *GroupUsecase) GetChatsAndUsersLocation(groupID string) (domain.GetChatsAndUsersLocationResponse, error) {
	ret := _m.Called(groupID)

	var r0 domain.GetChatsAndUsersLocationResponse
	if rf, ok := ret.Get(0).(func(string) domain.GetChatsAndUsersLocationResponse); ok {
		r0 = rf(groupID)
	} else {
		r0 = ret.Get(0).(domain.GetChatsAndUsersLocationResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupDetail provides a mock function with given fields: id
func (_m *GroupUsecase) GetGroupDetail(id string) (domain.Group, error) {
	ret := _m.Called(id)

	var r0 domain.Group
	if rf, ok := ret.Get(0).(func(string) domain.Group); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Group)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFiles provides a mock function with given fields: _a0, bucket, profileImg, id_group
func (_m *GroupUsecase) UploadFiles(_a0 *session.Session, bucket string, profileImg *multipart.FileHeader, id_group string) (string, error) {
	ret := _m.Called(_a0, bucket, profileImg, id_group)

	var r0 string
	if rf, ok := ret.Get(0).(func(*session.Session, string, *multipart.FileHeader, string) string); ok {
		r0 = rf(_a0, bucket, profileImg, id_group)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*session.Session, string, *multipart.FileHeader, string) error); ok {
		r1 = rf(_a0, bucket, profileImg, id_group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGroupUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroupUsecase creates a new instance of GroupUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroupUsecase(t mockConstructorTestingTNewGroupUsecase) *GroupUsecase {
	mock := &GroupUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
