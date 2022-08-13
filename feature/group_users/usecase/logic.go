package usecase

import (
	"errors"
	"lesgoobackend/domain"
)

type groupUsersData struct {
	groupUsersData domain.Group_UserData
}

// // LeaveGroup implements domain.Group_UserUsecase
func (gud *groupUsersData) LeaveGroup(data domain.Group_User) error {
	if data.Group_ID == "" || data.User_ID == 0 {
		return errors.New("failed")
	}

	err := gud.groupUsersData.Leave(data)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

// AddJoined implements domain.Group_UserUsecase
func (gud *groupUsersData) AddJoined(data domain.Group_User) error {
	if data.Group_ID == "" || data.Latitude == 0 || data.Longitude == 0 {
		return errors.New("failed")
	}

	err := gud.groupUsersData.Joined(data)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

func New(data domain.Group_UserData) domain.Group_UserUsecase {
	return &groupUsersData{
		groupUsersData: data,
	}
}
