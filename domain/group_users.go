package domain

type Group_User struct {
	ID        uint
	Group_ID  string
	User_ID   uint
	Longitude float64
	Latitude  float64
}
type Group_UserUsecase interface {
	AddJoined(data Group_User) error
	LeaveGroup(data Group_User) error
}

type Group_UserData interface {
	Joined(data Group_User) error
	Leave(data Group_User) error
}
