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
