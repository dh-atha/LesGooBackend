package data

import (
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID                 string `gorm:"type:VARCHAR(255);primaryKey"`
	Created_By_User_ID uint
	Name               string
	Description        string
	Start_Date         string
	End_Date           string
	Start_Dest         string
	Final_Dest         string
	GroupImg           string
	Status             string
	Group_Users        []groupUsersData.Group_User `gorm:"foreignKey:Group_ID"`
	Chats              []chatsData.Chat            `gorm:"foreignKey:Group_ID"`
}
