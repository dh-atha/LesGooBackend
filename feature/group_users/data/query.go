package data

import (
	"errors"
	"lesgoobackend/domain"
	"log"

	"gorm.io/gorm"
)

type groupUsersData struct {
	db *gorm.DB
}

// Leave implements domain.Group_UserData
func (gu *groupUsersData) Leave(data domain.Group_User) error {
	cnv := fromModelLeave(data)

	log.Println("cnv.Group_ID:", cnv.Group_ID, "cnv.User_ID:", cnv.User_ID)
	result := gu.db.Where("group_id = ? AND user_id = ?", cnv.Group_ID, cnv.User_ID).Delete(&Group_User{})
	if result.RowsAffected == 0 || result.Error != nil {
		return errors.New("cannot delete from group")
	}

	return nil
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
