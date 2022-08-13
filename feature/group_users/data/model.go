package data

import (
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type Group_User struct {
	gorm.Model
	Group_ID  string
	User_ID   uint
	Longitude float64
	Latitude  float64
}

func fromModelJoin(data domain.Group_User) Group_User {
	return Group_User{
		Group_ID:  data.Group_ID,
		User_ID:   data.User_ID,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
	}
}

func fromModelLeave(data domain.Group_User) Group_User {
	return Group_User{
		User_ID:  data.User_ID,
		Group_ID: data.Group_ID,
	}
}
