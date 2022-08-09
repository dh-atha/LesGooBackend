package domain

import "time"

type Group struct {
	ID                 string
	Created_By_User_ID uint
	Name               string
	Description        string
	Start_Dest         string
	Final_Dest         string
	GroupImg           string
	Created_At         time.Time
}

type GroupUsecase interface{}

type GroupData interface{}
