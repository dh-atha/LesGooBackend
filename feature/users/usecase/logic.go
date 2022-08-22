package usecase

import (
	"errors"
	"fmt"
	"lesgoobackend/domain"
	"lesgoobackend/infrastructure/aws/s3"
	"log"

	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUsecase {
	return &userUsecase{
		userData: ud,
		validate: v,
	}
}

func (ud *userUsecase) AddUser(newUser domain.User) (row int, err error) {
	checkDuplicate := ud.userData.CheckDuplicate(newUser)
	if checkDuplicate {
		return 0, errors.New("username or email already registered")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashed)
	user, err := ud.userData.Insert(newUser)
	return user, err
}

func (ud *userUsecase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)
	return response, data, err
}

func (ud *userUsecase) UpdateUser(id int, updateProfile domain.User) (row int, err error) {
	data, err := ud.userData.Update(id, updateProfile)
	return data, err
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
			return row, errors.New("record not found")
		} else {
			return row, errors.New("failed to delete user")
		}
	}
	return row, nil
}

func (ud *userUsecase) UploadFiles(session *session.Session, bucket string, profileImg *multipart.FileHeader) (string, error) {
	log.Println(bucket)
	profileImgExt := strings.Split(profileImg.Filename, ".")
	ext := profileImgExt[len(profileImgExt)-1]
	if ext != "png" && ext != "PNG" && ext != "jpeg" && ext != "JPEG" && ext != "jpg" && ext != "JPG" {
		return "", errors.New("image not supported, supported: png/jpeg/jpg")
	}

	destination := fmt.Sprint("images/", uuid.NewString(), "_", profileImg.Filename)
	profileImgUrl, err := s3.DoUpload(session, *profileImg, bucket, destination)
	if err != nil {
		return "", errors.New("cant upload image to s3")
	}

	return profileImgUrl, nil
}

func (ud *userUsecase) Logout(userID uint) error {
	err := ud.userData.Logout(userID)
	return err
}

func (ud *userUsecase) GetGroupID(data domain.User) string {
	groupID := ud.userData.GetGroupID(data)
	return groupID
}
