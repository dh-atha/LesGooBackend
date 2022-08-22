package usecase

import (
	"errors"
	"fmt"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/infrastructure/aws/s3"
	"lesgoobackend/mocks"
	"mime/multipart"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	usecase := New(repo, validator.New())
	insertData := domain.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@min.com",
		Phone:    "08123456789",
		Password: "12345678",
	}

	t.Run("duplicate data", func(t *testing.T) {
		repo.On("CheckDuplicate", insertData).Return(true, errors.New("Invalid Username")).Once()
		data, err := usecase.AddUser(insertData)
		assert.Equal(t, 0, data)
		assert.EqualError(t, err, "Invalid Username")
		repo.AssertExpectations(t)
	})

	t.Run("success add user", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false, nil).Once()
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		data, err := usecase.AddUser(insertData)
		assert.Equal(t, 1, data)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@min.com",
		Phone:    "08123456789",
		Password: "12345678",
	}
	outputData := domain.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@min.com",
		Phone:    "08123456789",
		Password: "12345678",
	}

	t.Run("Login Success", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(1, insertData, nil).Once()

		useCase := New(repo, validator.New())

		row, res, err := useCase.LoginUser(insertData)
		assert.Nil(t, err)
		assert.Equal(t, outputData, res)
		assert.Equal(t, 1, row)
		repo.AssertExpectations(t)
	})

	t.Run("Username Not Found", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(0, domain.User{}, gorm.ErrRecordNotFound, nil).Once()

		useCase := New(repo, validator.New())

		row, res, err := useCase.LoginUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.Username)
		// assert.Equal(t, err, gorm.ErrRecordNotFound.Error())
		// assert.Nil(t, res)
		assert.Equal(t, 0, row)
		repo.AssertExpectations(t)
	})

	t.Run("Login Wrong Pass", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(0, domain.User{}, gorm.ErrRecordNotFound, nil).Once()

		useCase := New(repo, validator.New())

		row, res, err := useCase.LoginUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.Password)
		assert.Equal(t, 0, row)
		repo.AssertExpectations(t)
	})

}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	useCase := New(repo, validator.New())
	insertData := domain.User{
		ID:         1,
		ProfileImg: "aaa.jpg",
		Username:   "admin",
		Email:      "admin@min.com",
		Phone:      "08123456789",
	}

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", insertData).Return(true, errors.New("Invalid Username")).Once()
		data, err := useCase.UpdateUser(1, insertData)
		assert.EqualError(t, err, "Invalid Username")
		assert.Equal(t, 0, data)
		repo.AssertExpectations(t)
	})

	t.Run("success update", func(t *testing.T) {
		repo.On("CheckDuplicate", insertData).Return(false, nil).Once()
		repo.On("Update", 1, insertData).Return(1, nil).Once()
		data, err := useCase.UpdateUser(1, insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, data)
		repo.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{
		ID:         1,
		ProfileImg: "aaa.jpg",
		Username:   "admin",
		Email:      "admin@min.com",
		Phone:      "08123456789",
	}
	outputData := domain.User{
		ID:         1,
		ProfileImg: "aaa.jpg",
		Username:   "admin",
		Email:      "admin@min.com",
		Phone:      "08123456789",
	}
	t.Run("Get User Success", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(insertData, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetProfile(int(insertData.ID))
		assert.Nil(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(domain.User{}, gorm.ErrRecordNotFound).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetProfile(int(insertData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, domain.User{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("Server Error", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(domain.User{}, errors.New("server error")).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.GetProfile(int(insertData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, domain.User{}, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{
		ID:         1,
		Username:   "admin",
		Email:      "admin@min.com",
		Password:   "asd123",
		Phone:      "08123456789",
		ProfileImg: "aaa.jpg",
		Fcm_Token:  "asdasdasd",
	}
	t.Run("Delete User Success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.DeleteUser(int(insertData.ID))
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})
	t.Run("Delete User Failed", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, fmt.Errorf("failed to delete user")).Once()

		useCase := New(repo, validator.New())

		_, err := useCase.DeleteUser(int(insertData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("failed to delete user"))
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, gorm.ErrRecordNotFound).Once()

		useCase := New(repo, validator.New())

		_, err := useCase.DeleteUser(int(insertData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
		repo.AssertExpectations(t)
	})
}

func TestUploadFiles(t *testing.T) {
	config := config.GetConfig()
	session := s3.ConnectAws(config)

	repo := mocks.UserData{}
	usecase := New(&repo, validator.New())

	imageFalse, _ := os.Open("./files/ERD.pdf")
	imageFalseCnv := &multipart.FileHeader{
		Filename: imageFalse.Name(),
	}

	imageTrue, _ := os.Open("./files/ERD.jpg")
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	t.Run("image not supported", func(t *testing.T) {
		profileImgUrl, err := usecase.UploadFiles(session, "bucket", imageFalseCnv)
		assert.Equal(t, profileImgUrl, "")
		assert.EqualError(t, err, "image not supported, supported: png/jpeg/jpg")
		repo.AssertExpectations(t)
	})

	t.Run("failed upload image", func(t *testing.T) {
		profileImgUrl, err := usecase.UploadFiles(session, "bucket", imageTrueCnv)
		assert.Equal(t, profileImgUrl, "")
		assert.EqualError(t, err, "cant upload image to s3")
		repo.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("success logout", func(t *testing.T) {
		repo.On("Logout", uint(1)).Return(nil).Once()

		usecase := New(repo, validator.New())
		err := usecase.Logout(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetGroupID(t *testing.T) {
	repo := new(mocks.UserData)
	usecase := New(repo, validator.New())
	insertData := domain.User{}

	t.Run("success get group id", func(t *testing.T) {
		repo.On("GetGroupID", mock.Anything).Return("m4nt4p").Once()
		groupID := usecase.GetGroupID(insertData)
		assert.Equal(t, groupID, "m4nt4p")
		repo.AssertExpectations(t)
	})
}
