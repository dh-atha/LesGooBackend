package usecase

import (
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
