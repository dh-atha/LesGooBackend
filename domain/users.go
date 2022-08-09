package domain

type User struct {
	ID         uint
	Username   string
	Email      string
	Password   string
	Phone      string
	ProfileImg string
}

type UserUsecase interface{}

type UserData interface{}
