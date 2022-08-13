package delivery

import "lesgoobackend/domain"

type GetChatsAndUsersLocationRequest struct {
	Group_ID string `json:"group_id" form:"group_id" validate:"required"`
}

type Group struct {
	ID                 string  `json:"id" form:"id" validate:"required"`
	Created_By_User_ID uint    `json:"created_by_user_id" form:"created_by_user_id"`
	Name               string  `json:"name" form:"name" validate:"required"`
	Description        string  `json:"description" form:"description" validate:"required"`
	Start_Dest         string  `json:"start_dest" form:"start_dest" validate:"required"`
	Final_Dest         string  `json:"final_dest" form:"final_dest" validate:"required"`
	Start_Date         string  `json:"start_date" form:"start_date" validate:"required"`
	End_Date           string  `json:"end_date" form:"end_date" validate:"required"`
	GroupImg           string  `json:"groupimg" form:"groupimg" validate:"required"`
	Longitude          float64 `json:"longitude" form:"longitude" validate:"required"`
	Latitude           float64 `json:"latitude" form:"latitude" validate:"required"`
}

type GroupUser struct {
	UserID    uint    `json:"user_id" form:"user_id" validate:"required"`
	GroupID   string  `json:"group_id" form:"group_id" validate:"required"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required"`
}

func ToModelGroup(data Group) domain.Group {
	return domain.Group{
		Created_By_User_ID: data.Created_By_User_ID,
		ID:                 data.ID,
		Name:               data.Name,
		Description:        data.Description,
		Start_Date:         data.Start_Date,
		End_Date:           data.End_Date,
		Start_Dest:         data.Start_Dest,
		Final_Dest:         data.Final_Dest,
		GroupImg:           data.GroupImg,
		Longitude:          data.Longitude,
		Latitude:           data.Latitude,
	}
}

func ToModelGroupUser(data GroupUser) domain.Group_User {
	return domain.Group_User{
		Group_ID:  data.GroupID,
		User_ID:   data.UserID,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
	}
}
