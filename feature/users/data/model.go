package data

import (
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"
	groupData "lesgoobackend/feature/groups/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	Phone       string
	ProfileImg  string
	Groups      []groupData.Group           `gorm:"foreignKey:Created_By_User_ID"`
	Group_Users []groupUsersData.Group_User `gorm:"foreignKey:User_ID"`
	Chats       []chatsData.Chat            `gorm:"foreignKey:User_ID"`
}
