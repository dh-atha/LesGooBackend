package usecase

import (
	"errors"
	"fmt"
	"lesgoobackend/domain"
	"lesgoobackend/infrastructure/aws/s3"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
)

type groupUsecase struct {
	groupData domain.GroupData
}

// UploadFiles implements domain.GroupUsecase
func (gu *groupUsecase) UploadFiles(session *session.Session, bucket string, groupImg *multipart.FileHeader, id_group string) (string, error) {
	groupImgExt := strings.Split(groupImg.Filename, ".")
	ext := groupImgExt[len(groupImgExt)-1]
	if ext != "png" && ext != "PNG" && ext != "jpeg" && ext != "JPEG" && ext != "jpg" && ext != "JPG" {
		return "", errors.New("image not supported, supported: png/jpeg/jpg")
	}

	destination := fmt.Sprint("images/", id_group, "_", groupImg.Filename)
	profileImgUrl, err := s3.DoUpload(session, *groupImg, bucket, destination)
	if err != nil {
		return "", errors.New("cant upload image to s3")
	}

	return profileImgUrl, nil
}

func (gu *groupUsecase) GetChatsAndUsersLocation(groupID string) (domain.GetChatsAndUsersLocationResponse, error) {
	res, err := gu.groupData.GetChatsAndUsersLocation(groupID)
	return res, err
}

// DeleteGroupByID implements domain.GroupUsecase
func (gu *groupUsecase) DeleteGroupByID(id_user uint) error {
	err := gu.groupData.RemoveGroupByID(id_user)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

// GetGroupDetail implements domain.GroupUsecase
func (gu *groupUsecase) GetGroupDetail(id string) (domain.Group, error) {
	response, err := gu.groupData.SelectSpecific(id)
	if err != nil {
		return domain.Group{}, errors.New("failed")
	}

	responseUser, errUser := gu.groupData.SelectUserData(response.ID)
	response.UsersbyID = responseUser
	if errUser != nil {
		return domain.Group{}, errors.New("failed")
	}

	return response, nil
}

// AddGroupUser implements domain.GroupUsecase
func (gu *groupUsecase) AddGroupUser(dataUser domain.Group_User) error {
	if dataUser.Group_ID == "" || dataUser.User_ID == 0 || dataUser.Latitude == 0 || dataUser.Longitude == 0 {
		return errors.New("failed")
	}

	err := gu.groupData.InsertGroupUser(dataUser)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

// AddGroup implements domain.GroupUsecase
func (gu *groupUsecase) AddGroup(dataGroup domain.Group) error {
	if dataGroup.Name == "" || dataGroup.Description == "" || dataGroup.Start_Dest == "" || dataGroup.Final_Dest == "" || dataGroup.GroupImg == "" {
		return errors.New("failed")
	}

	err := gu.groupData.InsertGroup(dataGroup)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

func New(gd domain.GroupData) domain.GroupUsecase {
	return &groupUsecase{
		groupData: gd,
	}
}
