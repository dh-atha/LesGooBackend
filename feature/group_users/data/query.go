package data

import (
	"errors"
	"lesgoobackend/domain"
	"log"
	"time"

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
	// verify that user not joined any group yet
	var groupUsersData Group_User
	res := gu.db.Where("user_id = ?", newJoined.User_ID).Find(&groupUsersData)
	if res.RowsAffected != 0 {
		return errors.New("cant join group when you are in a group")
	}

	var groupID string
	gu.db.Raw("SELECT id FROM groups WHERE id = ? AND deleted_at IS NULL", newJoined.Group_ID).Scan(&groupID)
	if groupID == "" {
		return errors.New("group not found")
	}

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

func (gud *groupUsersData) Update(data domain.Group_User) error {
	var get Group_User
	err := gud.db.Where("user_id = ? AND group_id = ?", data.User_ID, data.Group_ID).First(&get).Error
	if err != nil {
		return err
	}
	err = gud.db.Model(&Group_User{}).Where("id = ?", get.ID).Updates(&data).Update("updated_at", time.Now()).Error
	return err
}

func (gud *groupUsersData) GetToken(groupID string, userID uint) []string {
	var res []string
	gud.db.Raw("SELECT fcm_token FROM users WHERE id != ? AND id in (SELECT user_id FROM group_users WHERE deleted_at is null AND group_id = ?);", userID, groupID).Scan(&res)
	return res
}
