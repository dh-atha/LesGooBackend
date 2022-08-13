package delivery

import "lesgoobackend/domain"

type GroupUsers struct {
	UserID    uint    `json:"user_id" form:"user_id"`
	GroupID   string  `json:"group_id" form:"group_id" validate:"required"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required"`
}

func ToModelJoin(data GroupUsers) domain.Group_User {
	return domain.Group_User{
		User_ID:   data.UserID,
		Group_ID:  data.GroupID,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
	}
}

func ToModeLeave(data GroupUsers) domain.Group_User {
	return domain.Group_User{
		User_ID:  data.UserID,
		Group_ID: data.GroupID,
	}
}
