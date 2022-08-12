package data

import (
	"errors"
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type groupUsersData struct {
	db *gorm.DB
}

// JoinGroupByID implements domain.Group_UserData
func (gu *groupUsersData) Joined(newJoined domain.Group_User) error {
	cnv := fromModelJoin(newJoined)
	result := gu.db.Create(&cnv)
	if result.Error != nil {
		return errors.New("all data required must be filled")
	}

	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}

	return nil

}

func New(db *gorm.DB) domain.Group_UserData {
	return &groupUsersData{
		db: db,
	}
}
