package data

import "gorm.io/gorm"

type Group_User struct {
	gorm.Model
	Group_ID  string
	User_ID   uint
	Longitude string
	Latitude  string
}
