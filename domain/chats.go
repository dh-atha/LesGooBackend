package domain

import (
	"context"
	"time"

	"firebase.google.com/go/messaging"
)

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
	SendNotification(Chat, *messaging.Client, context.Context) (int, error)
}

type ChatData interface {
	Insert(Chat) error
	GetToken(groupID string, userID uint) []string
}
