package data

import (
	"errors"
	"fmt"
	"lesgoobackend/domain"
	"log"

	"golang.org/x/crypto/bcrypt"
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
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed insert data")
	}
	return int(result.RowsAffected), nil
}

func (ud *userData) Login(userLogin domain.User) (row int, data domain.User, err error) {
	var dataUser = FromModel(userLogin)
	password := dataUser.Password

	result := ud.db.Where("username = ?", dataUser.Username).First(&dataUser)

	if result.RowsAffected != 1 {
		return -1, domain.User{}, fmt.Errorf("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))

	if err != nil {
		return -2, domain.User{}, fmt.Errorf("wrong password")
	}

	if dataUser.Fcm_Token != "" {
		return -3, domain.User{}, errors.New("must logout from another device")
	} else {
		ud.db.Model(&dataUser).Update("fcm_token", userLogin.Fcm_Token)
	}

	return int(result.RowsAffected), dataUser.ToModel(), nil
}

func (ud *userData) Update(userID int, updatedData domain.User) (row int, err error) {
	var cnv = FromModel(updatedData)
	cnv.ID = uint(userID)
	result := ud.db.Model(&User{}).Where("ID = ?", userID).Updates(cnv)
	if result.Error != nil {
		log.Println("Cannot update data", errors.New("error db"))
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return -2, errors.New("failed update data")
	}

	return int(result.RowsAffected), nil
}

func (ud *userData) GetSpecific(userID int) (domain.User, error) {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}, errors.New("data not found")
	}

	return tmp.ToModel(), nil
}

func (ud *userData) Delete(userID int) (row int, err error) {
	res := ud.db.Delete(&User{}, userID)
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return 0, res.Error
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return 0, fmt.Errorf("failed to delete user")
	}
	return int(res.RowsAffected), nil
}

func (ud *userData) Logout(userID uint) error {
	err := ud.db.Model(&User{}).Where("id = ?", userID).Update("fcm_token", "").Error
	if err != nil {
		return err
	}
	return nil
}

func (ud *userData) CheckDuplicate(newUser domain.User) (bool, error) {
	var user = FromModel(newUser)
	err := ud.db.Find(&user, "username = ?", newUser.Username)
	if err.RowsAffected >= 1 {
		log.Println("Invalid Username", err.Error)
		return true, errors.New("Invalid Username")
	}
	err = ud.db.Find(&user, "email = ?", newUser.Email)
	if err.RowsAffected >= 1 {
		log.Println("Invalid Email", err.Error)
		return true, errors.New("Invalid Email")
	}
	err = ud.db.Find(&user, "phone = ?", newUser.Phone)
	if err.RowsAffected >= 1 {
		log.Println("Invalid Phone", err.Error)
		return true, errors.New("Invalid Phone")
	}

	return false, nil
}

func (ud *userData) GetGroupID(data domain.User) string {
	var result string
	ud.db.Raw("SELECT group_id FROM group_users WHERE deleted_at is NULL AND user_id = ?", data.ID).Scan(&result)
	return result
}
