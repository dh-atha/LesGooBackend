package usecase

import (
	"errors"
	"fmt"
	"lesgoobackend/domain"
	"lesgoobackend/mocks"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@min.com",
		Phone:    "08123456789",
		Password: "12345678",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddUser(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated Data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, gorm.ErrRegistered).Once()

		useCase := New(repo, validator.New())

		row, err := useCase.AddUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, -4, row)
		repo.AssertExpectations(t)
	})

	t.Run("Error from server", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, gorm.ErrInvalidValueOfLength).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddUser(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, -4, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Username", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Username = ""
		res, err := useCase.AddUser(dummy)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("invalid username").Error())
		assert.Equal(t, -1, res)
	})

	t.Run("Empty Password", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Password = ""
		res, err := useCase.AddUser(dummy)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("invalid password").Error())
		assert.Equal(t, -1, res)
	})

	t.Run("Empty Phone", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Phone = ""
		res, err := useCase.AddUser(dummy)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("invalid phone number").Error())
		assert.Equal(t, -1, res)
	})

	t.Run("Empty Email", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.Email = ""
		res, err := useCase.AddUser(dummy)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("invalid Email").Error())
		assert.Equal(t, -1, res)
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
	insertData := domain.User{
		ID:         1,
		ProfileImg: "aaa.jpg",
		Username:   "admin",
		Email:      "admin@min.com",
		Phone:      "08123456789",
	}
	// outputData := domain.User{
	// 	ID:       1,
	// 	Fullname: "Rizuana Nadifatul",
	// 	Username: "rizunadiva",
	// 	Phone:    "081936665965",
	// 	Password: "12345678",
	// }
	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.UpdateUser(int(insertData.ID), insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})
	t.Run("Username or phone number already exist", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(0, errors.New("username or phone number already exist")).Once()

		useCase := New(repo, validator.New())

		_, err := useCase.UpdateUser(int(insertData.ID), insertData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.New("username or phone number already exist").Error())
		repo.AssertExpectations(t)
	})
	t.Run("Error from server", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(0, gorm.ErrInvalidValueOfLength).Once()

		useCase := New(repo, validator.New())

		insertData.Username = "123aoeijakdngnsvbsnzoczbjfakdjfoadijfoangnbcoloijapdfaposdjfpk"
		res, err := useCase.UpdateUser(int(insertData.ID), insertData)
		assert.NotNil(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty ProfileImg", func(t *testing.T) {})
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
	t.Run("Get User Failed", func(t *testing.T) {
		repo.On("GetSpecific", mock.Anything).Return(domain.User{}, gorm.ErrRecordNotFound).Once()

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
}

// func TestUploadFiles(t *testing.T) {
// 	imageTrue, _ := os.Open("./files/ERD.jpg")
// 	imageTrueCnv := &multipart.FileHeader{
// 		Filename: imageTrue.Name(),
// 		Size:     int64(imageTrue.Fd()),
// 	}

// 	config := config.GetConfig()
// 	session := s3.ConnectAws(config)

// 	repo := mocks.UserData{}
// 	usecase := New(&repo, validator.New())

// 	t.Run("cant upload image to s3", func(t *testing.T) {
// 		repo.On("UploadFiles", mock.Anything, mock.Anything).Return(nil, errors.New("cant upload image to s3")).Once()
// 		profileImgUrl, err := usecase.UploadFiles(session, "bucket", imageTrueCnv)

// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", profileImgUrl)
// 		repo.AssertExpectations(t)
// 	})
// }

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
