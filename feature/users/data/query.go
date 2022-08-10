package data

import (
	"errors"
	"lesgoobackend/domain"
	"log"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.UserData {
	return &userData{
		db: DB,
	}
}

func (ud *userData) Insert(newUser domain.User) (row int, err error) {
	var cnv = FromModel(newUser)
	result := ud.db.Create(&cnv)
	if result.Error != nil {
		log.Println("Cannot create object", errors.New("error db"))
		return -1, errors.New("username or number phone already exist")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}
