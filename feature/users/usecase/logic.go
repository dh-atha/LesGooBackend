package usecase

import (
	"errors"
	"lesgoobackend/domain"
	"lesgoobackend/feature/users/delivery"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (ud *userUsecase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)

	return response, data, err
}

func (ud *userUsecase) UpdateUser(id int, updateProfile domain.User) (row int, err error) {
	var tmp delivery.UpdateFormat
	qry := map[string]interface{}{}
	if tmp.Username != "" {
		qry["username"] = &tmp.Username
	}
	if tmp.Email != "" {
		qry["email"] = &tmp.Email
	}
	if tmp.Phone != "" {
		qry["phone"] = &tmp.Phone
	}

	if id == -1 {
		return 0, errors.New("invalid user")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(updateProfile.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrypt password", err)
		return 0, err
	}

	updateProfile.Password = string(hashed)
	result, err := ud.userData.Update(id, updateProfile)
	if err != nil {
		return 0, errors.New("username or phone number already exist")
	}

	return result, nil
}

func (ud *userUsecase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (ud *userUsecase) DeleteUser(id int) (row int, err error) {
	row, err = ud.userData.Delete(id)
	if err != nil {
		log.Println("delete usecase error", err.Error())
		if err == gorm.ErrRecordNotFound {
			return row, errors.New("data not found")
		} else {
			return row, errors.New("failed to delete user")
		}
	}
	return row, nil
}
