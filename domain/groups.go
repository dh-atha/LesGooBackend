package domain

import "time"

type Group struct {
	ID                 string
	Created_By_User_ID uint
	Name               string
	Description        string
	Start_Date         string
	End_Date           string
	Start_Dest         string
	Final_Dest         string
	GroupImg           string
	Status             string
	Created_At         time.Time
}

type GroupUsecase interface {
	// AddGroup() // Add jadiin statusnya active
	// GetGroupDetail()
	// JoinGroupByID()
	// DeleteGroupByID() // Delete jadiin statusnya inactive
	GetChatsAndUsersLocation(groupID string) (GetChatsAndUsersLocationResponse, error)
	// LeaveGroup()
}

type GroupData interface {
	// Insert()
	// GetSpecific()
	// JoinGroupByID()
	// Delete()
	GetChatsAndUsersLocation(groupID string) (GetChatsAndUsersLocationResponse, error)
	// Leave()
}

type GetChatsAndUsersLocationResponse struct {
	Group_ID    string                    `json:"group_id" form:"group_id"`
	Name        string                    `json:"name" form:"name"`
	Chats       []JoinChatsWithUsers      `json:"chats" form:"chats"`
	Group_Users []JoinGroupUsersWithUsers `json:"group_users" form:"group_users"`
}

type JoinChatsWithUsers struct {
	ID         uint      `json:"id" form:"id"`
	Message    string    `json:"message" form:"message"`
	User_ID    uint      `json:"user_id" form:"user_id"`
	ProfileImg string    `json:"profileimg" form:"profileimg"`
	Username   string    `json:"username" form:"username"`
	Created_At time.Time `json:"created_at" form:"created_at"`
}

type JoinGroupUsersWithUsers struct {
	ID         uint    `json:"id" form:"id"`
	Latitude   float64 `json:"latitude" form:"latitude"`
	Longitude  float64 `json:"longitude" form:"longitude"`
	User_ID    uint    `json:"user_id" form:"user_id"`
	Username   string  `json:"username" form:"username"`
	ProfileImg string  `json:"profileimg" form:"profileimg"`
}
