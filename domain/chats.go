package domain

import "time"

type Chat struct {
	ID         uint
	Group_ID   string
	User_ID    uint
	Message    string
	IsSOS      bool
	Created_At time.Time
}

type ChatUsecase interface {
	SendChats(Chat) error
	SendNotification(Chat) (int, error)
}

type ChatData interface {
	Insert(Chat) error
	GetToken(groupID string, userID uint) []string
}
