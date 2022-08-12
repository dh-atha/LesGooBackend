package usecase

import (
	"errors"
	"lesgoobackend/domain"
)

type groupUsecase struct {
	groupData domain.GroupData
}

func New(gd domain.GroupData) domain.GroupUsecase {
	return &groupUsecase{
		groupData: gd,
	}
}

func (gu *groupUsecase) GetChatsAndUsersLocation(groupID string) (domain.GetChatsAndUsersLocationResponse, error) {
	res, err := gu.groupData.GetChatsAndUsersLocation(groupID)
	return res, err
}

// AddGroupUser implements domain.GroupUsecase
func (gu *groupUsecase) AddGroupUser(dataUser domain.Group_User) error {
	if dataUser.Group_ID == "" || dataUser.User_ID == 0 || dataUser.Latitude == 0 || dataUser.Longitude == 0 {
		return errors.New("failed")
	}

	err := gu.groupData.InsertGroupUser(dataUser)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

// AddGroup implements domain.GroupUsecase
func (gu *groupUsecase) AddGroup(dataGroup domain.Group) error {
	if dataGroup.Name == "" || dataGroup.Description == "" || dataGroup.Start_Dest == "" || dataGroup.Final_Dest == "" || dataGroup.Latitude == 0 || dataGroup.Longitude == 0 {
		return errors.New("failed")
	}

	err := gu.groupData.InsertGroup(dataGroup)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}
