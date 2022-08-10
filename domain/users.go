package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID         uint
	Username   string
	Email      string
	Password   string
	Phone      string
	ProfileImg string
}

type UserHandler interface {
	InsertUser() echo.HandlerFunc
	LoginHandler() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UserUsecase interface{
	AddUser(newUser User) (row int, err error)
}

type UserData interface{
	Insert(newUser User) (row int, err error)
}
