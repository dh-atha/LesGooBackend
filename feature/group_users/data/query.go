package data

import (
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type groupUsersData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Group_UserData {
	return &groupUsersData{
		db: db,
	}
}
