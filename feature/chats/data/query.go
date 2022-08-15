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

func (cd *chatData) Insert(data domain.Chat) error {
	var convert = ToEntity(data)
	err := cd.db.Create(&convert).Error
	if err != nil {
		return err
	}
	return nil
}

func (cd *chatData) GetToken(groupID string, userID uint) []string {
	var res []string
	cd.db.Raw("SELECT fcm_token FROM users WHERE id in (SELECT user_id FROM group_users WHERE group_id = ? AND user_id != ?)", groupID, userID).Scan(&res)
	return res
}
