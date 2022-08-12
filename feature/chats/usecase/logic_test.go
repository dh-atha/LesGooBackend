package usecase

import (
	"lesgoobackend/domain"
	"lesgoobackend/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendChat(t *testing.T) {
	repo := mocks.ChatData{}
	usecase := New(&repo)
	successSend := domain.Chat{Group_ID: "m4nt4p", User_ID: 1, Message: "Hello"}

	t.Run("success add to cart", func(t *testing.T) {
		repo.On("Insert", successSend).Return(nil).Once()

		err := usecase.SendChats(successSend)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
