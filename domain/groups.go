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
	// GetChatsAndUsersLocation()
	// LeaveGroup()
}

type GroupData interface {
	// Insert()
	// GetSpecific()
	// JoinGroupByID()
	// Delete()
	// GetChatsAndUsersLocation()
	// Leave()
}
