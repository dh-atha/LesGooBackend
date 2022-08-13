package usecase

import (
	"errors"
	"lesgoobackend/domain"
	"lesgoobackend/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
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
