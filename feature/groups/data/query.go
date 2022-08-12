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

func (gd *groupData) GetChatsAndUsersLocation(groupID string) (domain.GetChatsAndUsersLocationResponse, error) {
	var result domain.GetChatsAndUsersLocationResponse
	result.Group_ID = groupID
	gd.db.Raw("SELECT id, name FROM groups WHERE id = ?", groupID).Scan(&result)
	gd.db.Raw("SELECT c.id, c.message, c.user_id, u.profile_img, u.username, c.created_at FROM chats c INNER JOIN users u ON u.id = c.user_id").Scan(&result.Chats)
	gd.db.Raw("SELECT g.id, g.latitude, g.longitude, g.user_id, g.user_id, u.username, u.profile_img FROM group_users g INNER JOIN users u ON u.id = g.user_id").Scan(&result.Group_Users)
	return result, nil
}
