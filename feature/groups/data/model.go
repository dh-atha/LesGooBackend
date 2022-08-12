package data

import (
	"lesgoobackend/domain"
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Group_ID           string `gorm:"type:VARCHAR(255);primaryKey"`
	Created_By_User_ID uint
	Name               string
	Description        string
	Start_Date         string
	End_Date           string
	Start_Dest         string
	Final_Dest         string
	GroupImg           string
	Status             string
	Longitude          float64
	Latitude           float64
	Group_Users        []groupUsersData.Group_User `gorm:"foreignKey:Group_ID"`
	Chats              []chatsData.Chat            `gorm:"foreignKey:Group_ID"`
}

type Group_User struct {
	gorm.Model
	Group_ID  string
	User_ID   uint
	Longitude float64
	Latitude  float64
}

func fromModelGroup(data domain.Group) Group {
	return Group{
		Model: gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Group_ID:           data.GroupID,
		Created_By_User_ID: data.Created_By_User_ID,
		Name:               data.Name,
		Description:        data.Description,
		Start_Date:         data.Start_Date,
		End_Date:           data.End_Date,
		Start_Dest:         data.Start_Dest,
		Final_Dest:         data.Final_Dest,
		GroupImg:           data.GroupImg,
		Status:             data.Status,
		Longitude:          data.Longitude,
		Latitude:           data.Latitude,
	}
}

func fromModelGroupUser(data domain.Group_User) Group_User {
	return Group_User{
		Model: gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Group_ID:  data.Group_ID,
		User_ID:   data.User_ID,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
	}
}
