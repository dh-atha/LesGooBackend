package usecase

import (
	"context"
	"errors"
	"lesgoobackend/domain"
	fcm "lesgoobackend/infrastructure/firebase/messaging"

	"firebase.google.com/go/messaging"
)

type groupUsersData struct {
	groupUsersData domain.Group_UserData
}

func New(data domain.Group_UserData) domain.Group_UserUsecase {
	return &groupUsersData{
		groupUsersData: data,
	}
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

func (gud *groupUsersData) UpdateLocation(data domain.Group_User, client *messaging.Client, context context.Context) error {
	err := gud.groupUsersData.Update(data)
	if err != nil {
		return err
	}
	res := gud.groupUsersData.GetToken(data.Group_ID, data.User_ID)
	fcm.SendLocation(res, client, context)
	return err
}
