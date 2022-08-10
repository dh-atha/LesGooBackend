package usecase

import (
	"errors"
	"lesgoobackend/domain"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData domain.UserData
}

func New(ud domain.UserData) domain.UserUsecase {
	return &userUsecase{
		userData: ud,
	}
}

func (ud *userUsecase) AddUser(newUser domain.User) (row int, err error) {

	if newUser.Username == "" {
		return -1, errors.New("invalid username")
	}
	if newUser.Email == "" {
		return -1, errors.New("invalid Email")
	}
	if newUser.Password == "" {
		return -1, errors.New("invalid password")
	}
	if newUser.Phone == "" {
		return -1, errors.New("invalid phone number")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrypt password", err)
		return -3, err
	}
	newUser.Password = string(hashed)
	inserted, err := ud.userData.Insert(newUser)

	if err != nil {
		log.Println("error from usecase", err.Error())
		return -4, err
	}
	return inserted, nil
}
