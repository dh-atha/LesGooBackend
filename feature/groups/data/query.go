package data

import (
	"errors"
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
	gd.db.Raw("SELECT id, name, status FROM groups WHERE id = ?", groupID).Scan(&result)
	gd.db.Raw("SELECT c.id, c.message, c.user_id, u.profile_img, u.username, c.created_at FROM chats c INNER JOIN users u ON u.id = c.user_id").Scan(&result.Chats)
	gd.db.Raw("SELECT g.id, g.latitude, g.longitude, g.user_id, g.user_id, u.username, u.profile_img FROM group_users g INNER JOIN users u ON u.id = g.user_id").Scan(&result.Group_Users)
	return result, nil
}

// RemoveGroupByID implements domain.GroupData
func (gd *groupData) RemoveGroupByID(id string, id_user uint) error {
	dataGroup := Group{}

	res := gd.db.Where("id = ? AND created_by_user_id = ?", id, id_user).Delete(&dataGroup)
	if res.RowsAffected == 0 || res.Error != nil {
		return res.Error
	}

	return nil

}

// SelectSpecific implements domain.GroupData
func (gd *groupData) SelectSpecific(id string) (domain.Group, error) {
	dataGroup := Group{}
	result := gd.db.Find(&dataGroup, id)
	if result.Error != nil {
		return domain.Group{}, result.Error
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

	result := gd.db.Create(&cnv)

	if result.Error != nil {
		return errors.New("all data required must be filled")
	}

	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}

	return nil
}
