package usecase

import (
	"errors"
	"lesgoobackend/domain"
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
