package data

import (
	"lesgoobackend/domain"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Group_ID string
	User_ID  uint
	Message  string
	IsSOS    bool
}

func ToEntity(data domain.Chat) Chat {
	return Chat{
		Group_ID: data.Group_ID,
		User_ID:  data.User_ID,
		Message:  data.Message,
		IsSOS:    data.IsSOS,
	}
}
