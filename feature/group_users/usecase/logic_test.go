package usecase

import (
	"context"
	"errors"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/infrastructure/firebase/messaging"
	"lesgoobackend/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddJoined(t *testing.T) {
	repo := mocks.Group_UserData{}
	usecase := New(&repo)
	successAdd := domain.Group_User{Group_ID: "m4nt4p", User_ID: 1, Longitude: 1.0, Latitude: 1.0}
	failedAdd := domain.Group_User{Group_ID: "", User_ID: 0, Longitude: 0.0, Latitude: 0.0}

	t.Run("success join", func(t *testing.T) {
		repo.On("Joined", successAdd).Return(nil).Once()
		err := usecase.AddJoined(successAdd)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed join", func(t *testing.T) {
		repo.On("Joined", failedAdd).Return(errors.New("failed")).Once()
		err := usecase.AddJoined(failedAdd)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestLeaveGroup(t *testing.T) {
	repo := mocks.Group_UserData{}
	usecase := New(&repo)
	successLeave := domain.Group_User{Group_ID: "m4nt4p", User_ID: 1, Longitude: 1.0, Latitude: 1.0}
	failedLeave := domain.Group_User{Group_ID: "", User_ID: 0, Longitude: 0.0, Latitude: 0.0}

	t.Run("success operation", func(t *testing.T) {
		repo.On("Leave", successLeave).Return(nil).Once()
		err := usecase.LeaveGroup(successLeave)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Leave", failedLeave).Return(errors.New("failed")).Once()
		err := usecase.LeaveGroup(failedLeave)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateLocation(t *testing.T) {
	repo := mocks.Group_UserData{}
	usecase := New(&repo)
	successLeave := domain.Group_User{Group_ID: "m4nt4p", User_ID: 1, Longitude: 1.0, Latitude: 1.0}
	ctx := context.Background()
	cfg := config.GetConfig()
	client := messaging.InitFirebaseClient(ctx, cfg)

	t.Run("Update error", func(t *testing.T) {
		repo.On("Update", successLeave).Return(errors.New("update error")).Once()
		err := usecase.UpdateLocation(successLeave, client, ctx)
		assert.EqualError(t, err, "update error")
		repo.AssertExpectations(t)
	})

	t.Run("Token not exists", func(t *testing.T) {
		repo.On("Update", successLeave).Return(nil).Once()
		repo.On("GetToken", successLeave.Group_ID, successLeave.User_ID).Return([]string{}).Once()
		err := usecase.UpdateLocation(successLeave, client, ctx)
		assert.EqualError(t, err, "no notification sent")
		repo.AssertExpectations(t)
	})

	t.Run("Token exists", func(t *testing.T) {
		repo.On("Update", successLeave).Return(nil).Once()
		repo.On("GetToken", successLeave.Group_ID, successLeave.User_ID).Return([]string{"asdasdasdasd"}).Once()
		err := usecase.UpdateLocation(successLeave, client, ctx)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
