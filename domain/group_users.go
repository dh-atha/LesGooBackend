package domain

import (
	"context"

	"firebase.google.com/go/messaging"
)

type Group_User struct {
	ID        uint
	Group_ID  string
	User_ID   uint
	Longitude float64
	Latitude  float64
}
type Group_UserUsecase interface {
	AddJoined(data Group_User) error
	LeaveGroup(data Group_User) error
	UpdateLocation(data Group_User, client *messaging.Client, context context.Context) error
}

type Group_UserData interface {
	Joined(data Group_User) error
	Leave(data Group_User) error
	Update(data Group_User) error
	GetToken(groupID string, userID uint) []string
}
