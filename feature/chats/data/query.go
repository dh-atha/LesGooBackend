package data

import (
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type chatData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ChatData {
	return &chatData{
		db: db,
	}
}
