package usecase

import (
	"lesgoobackend/domain"
)

type groupUsersData struct {
	groupUsersData domain.Group_UserData
}

func New(data domain.Group_UserData) domain.Group_UserUsecase {
	return &groupUsersData{
		groupUsersData: data,
	}
}
