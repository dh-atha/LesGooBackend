package domain

import (
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
)

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
	UsersbyID          []UsersbyID
}

type UsersbyID struct {
	UserID   uint
	Username string
}

type GroupUsecase interface {
	AddGroup(data Group) error
	AddGroupUser(dataUser Group_User) error
	GetGroupDetail(id string) (Group, error)
	DeleteGroupByID(groupID string, userID uint) error
	UploadFiles(session *session.Session, bucket string, profileImg *multipart.FileHeader, id_group string) (string, error)
	GetChatsAndUsersLocation(groupID string) (GetChatsAndUsersLocationResponse, error)
}

type GroupData interface {
	InsertGroup(data Group) error
	InsertGroupUser(dataUser Group_User) error
	SelectSpecific(id string) (Group, error)
	SelectUserData(id string) ([]UsersbyID, error)
	RemoveGroupByID(groupID string, userID uint) error
	GetChatsAndUsersLocation(groupID string) (GetChatsAndUsersLocationResponse, error)
}

type GetChatsAndUsersLocationResponse struct {
	Group_ID    string                    `json:"group_id" form:"group_id"`
	Name        string                    `json:"name" form:"name"`
	Status      string                    `json:"status" form:"status"`
	Start_Dest  string                    `json:"start_dest" form:"start_dest"`
	Final_Dest  string                    `json:"final_dest" form:"final_dest"`
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
