package data

import (
	"errors"
	"fmt"
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

func (gd *groupData) GetChatsAndUsersLocation(groupID string) (domain.GetChatsAndUsersLocationResponse, error) {
	var result domain.GetChatsAndUsersLocationResponse
	result.Group_ID = groupID
	gd.db.Raw("SELECT id, name, status, start_dest, final_dest FROM `groups` WHERE id = ?", groupID).Scan(&result)
	gd.db.Raw("SELECT c.id, c.message, c.user_id, u.profile_img, u.username, c.created_at FROM chats c INNER JOIN users u ON u.id = c.user_id WHERE c.is_sos = false AND c.deleted_at is NULL AND group_id = ?", groupID).Scan(&result.Chats)
	gd.db.Raw("SELECT g.id, g.latitude, g.longitude, g.user_id, g.user_id, u.username, u.profile_img FROM group_users g INNER JOIN users u ON u.id = g.user_id WHERE g.deleted_at is NULL AND group_id = ?", groupID).Scan(&result.Group_Users)
	return result, nil
}

// RemoveGroupByID implements domain.GroupData
func (gd *groupData) RemoveGroupByID(groupID string, userID uint) error {
	dataGroup := Group{}

	err := gd.db.Where("id = ?", groupID).First(&dataGroup).Error
	if err != nil {
		return err
	}

	if dataGroup.Created_By_User_ID != userID {
		return errors.New("unauthorized")
	}

	dataGroup.Status = "non-active"
	err = gd.db.Save(&dataGroup).Error
	if err != nil {
		return err
	}

	gd.db.Delete(&dataGroup)

	// Delete data from group_users table
	var groupUsersData Group_User
	gd.db.Model(&groupUsersData).Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&groupUsersData)

	return nil

}

// SelectSpecific implements domain.GroupData
func (gd *groupData) SelectSpecific(id string) (domain.Group, error) {
	dataGroup := Group{}
	result := gd.db.Find(&dataGroup, "id = ?", id)
	if result.Error != nil {
		return domain.Group{}, result.Error
	}

	if dataGroup.Name == "" {
		return domain.Group{}, errors.New("group not found")
	}

	return toDomainByID(dataGroup), nil
}

// SelectUserData implements domain.GroupData
func (gd *groupData) SelectUserData(id string) ([]domain.UsersbyID, error) {
	dataGroupUsers := []Group_User{}

	result := gd.db.Preload("User").Find(&dataGroupUsers, "group_id", id)
	if result.Error != nil {
		return []domain.UsersbyID{}, result.Error
	}

	return ToUsersDomainList(dataGroupUsers), nil
}

// InsertGroupUser implements domain.GroupData
func (gd *groupData) InsertGroupUser(newGroupUser domain.Group_User) error {

	cnv := fromModelGroupUser(newGroupUser)

	result := gd.db.Create(&cnv)

	if result.Error != nil {
		return errors.New("all data required must be filled")
	}

	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}

	return nil
}

func (gd *groupData) InsertGroup(newGroup domain.Group) error {

	//	Set status active
	newGroup.Status = "active"

	cnv := fromModelGroup(newGroup)

	var groupUsersData Group_User
	res := gd.db.Where("user_id = ?", newGroup.Created_By_User_ID).Find(&groupUsersData)
	if res.RowsAffected != 0 {
		return errors.New("cant create group when you are in a group")
	}

	result := gd.db.Create(&cnv)

	if result.Error != nil {
		return errors.New("all data required must be filled")
	}

	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}

	fmt.Println("RowsAffected:", result.RowsAffected)

	return nil
}
