package data

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Group_ID string
	User_ID  uint
	Message  string
	IsSOS    bool
}
