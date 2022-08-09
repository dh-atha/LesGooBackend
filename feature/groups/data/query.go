package data

import (
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type groupData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.GroupData {
	return &groupData{
		db: db,
	}
}
