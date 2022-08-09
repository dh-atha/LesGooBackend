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

type ChatUsecase interface{}

type ChatData interface{}
