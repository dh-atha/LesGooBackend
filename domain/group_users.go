package domain

type Group_User struct {
	ID        uint
	Group_ID  string
	User_ID   uint
	Longitude string
	Latitude  string
}

type Group_UserUsecase interface{}

type Group_UserData interface{}
