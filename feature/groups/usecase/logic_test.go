package usecase

import (
	"errors"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/infrastructure/aws/s3"
	"lesgoobackend/mocks"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetChatsAndUsersLocation(t *testing.T) {
	t.Run("success get chats and users location", func(t *testing.T) {
		repo := mocks.GroupData{}
		usecase := New(&repo)
		successGet := domain.GetChatsAndUsersLocationResponse{Group_ID: "m4nt4p", Name: "Udin"}
		repo.On("GetChatsAndUsersLocation", "m4nt4p").Return(successGet, nil).Once()

		response, err := usecase.GetChatsAndUsersLocation("m4nt4p")
		assert.Nil(t, err)
		assert.Equal(t, successGet, response)
		repo.AssertExpectations(t)
	})
}

func TestDeleteGroupByID(t *testing.T) {
	t.Run("success operation", func(t *testing.T) {
		repo := mocks.GroupData{}
		usecase := New(&repo)
		repo.On("RemoveGroupByID", "m4nt4p", uint(1)).Return(nil).Once()

		err := usecase.DeleteGroupByID("m4nt4p", uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed operation", func(t *testing.T) {
		repo := mocks.GroupData{}
		usecase := New(&repo)
		repo.On("RemoveGroupByID", "m4nt4p", uint(1)).Return(errors.New("failed")).Once()

		err := usecase.DeleteGroupByID("m4nt4p", uint(1))
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetGroupDetail(t *testing.T) {
	repo := mocks.GroupData{}
	usecase := New(&repo)
	failed := domain.Group{}
	success := domain.Group{ID: "m4nt4p"}

	t.Run("failed get groups", func(t *testing.T) {
		repo.On("SelectSpecific", "m4nt4p").Return(failed, errors.New("failed")).Once()
		response, err := usecase.GetGroupDetail("m4nt4p")
		assert.EqualError(t, err, "failed")
		assert.Equal(t, failed, response)
		repo.AssertExpectations(t)
	})

	t.Run("failed get userData", func(t *testing.T) {
		repo.On("SelectSpecific", "m4nt4p").Return(success, nil).Once()
		repo.On("SelectUserData", success.ID).Return([]domain.UsersbyID{}, errors.New("failed")).Once()
		response, err := usecase.GetGroupDetail("m4nt4p")
		assert.EqualError(t, err, "failed")
		assert.Equal(t, failed, response)
		repo.AssertExpectations(t)
	})

	t.Run("success operation", func(t *testing.T) {
		repo.On("SelectSpecific", success.ID).Return(success, nil).Once()
		repo.On("SelectUserData", success.ID).Return([]domain.UsersbyID{{UserID: 1, Username: "atha"}, {UserID: 2, Username: "faqih"}}, nil).Once()
		response, err := usecase.GetGroupDetail(success.ID)
		assert.Nil(t, err)
		assert.Equal(t, success.ID, response.ID)
		repo.AssertExpectations(t)
	})
}

//TestAddGroup implements domain.GroupUsecase
func TestAddGroup(t *testing.T) {
	repo := mocks.GroupData{}
	usecase := New(&repo)

	// successAdd := domain.Group{ID: "m4nt4p"}
	failedInsert := domain.Group{ID: "m4nt4p", Name: "Udin", Description: "Udin", Start_Dest: "asda", Final_Dest: "asada"}
	insertDB := domain.Group{ID: "m4nt4p", Name: "Udin", Description: "Udin", Start_Dest: "asda", Final_Dest: "asada", GroupImg: "sadas"}

	t.Run("failed at verification", func(t *testing.T) {
		err := usecase.AddGroup(failedInsert)
		assert.EqualError(t, err, "failed")
		repo.AssertExpectations(t)
	})

	t.Run("failed insert group", func(t *testing.T) {
		repo.On("InsertGroup", insertDB).Return(errors.New("error insert group")).Once()
		err := usecase.AddGroup(insertDB)
		assert.EqualError(t, err, "error insert group")
		repo.AssertExpectations(t)
	})

	t.Run("success insert group", func(t *testing.T) {
		repo.On("InsertGroup", insertDB).Return(nil).Once()
		err := usecase.AddGroup(insertDB)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestAddGroupUser(t *testing.T) {
	repo := mocks.GroupData{}
	usecase := New(&repo)
	failedVerif := domain.Group_User{Longitude: 1.232312}
	successVerif := domain.Group_User{Group_ID: "asda", User_ID: 1, Latitude: 1.12312, Longitude: 1.232312}

	t.Run("failed at verification input", func(t *testing.T) {
		err := usecase.AddGroupUser(failedVerif)
		assert.EqualError(t, err, "failed")
		repo.AssertExpectations(t)
	})

	t.Run("failed at insertDB", func(t *testing.T) {
		repo.On("InsertGroupUser", mock.Anything).Return(errors.New("ada errorrrrr di data layer")).Once()
		err := usecase.AddGroupUser(successVerif)
		assert.EqualError(t, err, "failed")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("InsertGroupUser", mock.Anything).Return(nil).Once()
		err := usecase.AddGroupUser(successVerif)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUploadFiles(t *testing.T) {
	config := config.GetConfig()
	session := s3.ConnectAws(config)

	repo := mocks.GroupData{}
	usecase := New(&repo)

	imageFalse, _ := os.Open("./files/ERD.pdf")
	imageFalseCnv := &multipart.FileHeader{
		Filename: imageFalse.Name(),
	}

	imageTrue, _ := os.Open("./files/ERD.jpg")
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	t.Run("image not supported", func(t *testing.T) {
		profileImgUrl, err := usecase.UploadFiles(session, "bucket", imageFalseCnv, "testgroup")
		assert.Equal(t, profileImgUrl, "")
		assert.EqualError(t, err, "image not supported, supported: png/jpeg/jpg")
		repo.AssertExpectations(t)
	})

	t.Run("failed upload image", func(t *testing.T) {
		profileImgUrl, err := usecase.UploadFiles(session, "bucket", imageTrueCnv, "testgroup")
		assert.Equal(t, profileImgUrl, "")
		assert.EqualError(t, err, "cant upload group image to s3")
		repo.AssertExpectations(t)
	})
}
