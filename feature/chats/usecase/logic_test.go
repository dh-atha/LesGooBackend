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

func TestSendNotification(t *testing.T) {
	repo := mocks.ChatData{}
	usecase := New(&repo)
	successSend := domain.Chat{Group_ID: "m4nt4p", User_ID: 1, Message: "Hello"}
	successGetUserData := domain.User{ID: 1, Username: "atha", Email: "email5atha@gmail.com", Password: "pass", Phone: "0822", ProfileImg: "utl.com", Fcm_Token: "yahahaha"}
	ctx := context.Background()
	cfg := config.GetConfig()
	client := messaging.InitFirebaseClient(ctx, cfg)

	t.Run("user not exists", func(t *testing.T) {
		repo.On("GetUserData", uint(1)).Return(domain.User{}, errors.New("record not found")).Once()
		data, err := usecase.SendNotification(successSend, client, ctx)
		assert.EqualError(t, err, "record not found")
		assert.Equal(t, 0, data)
		repo.AssertExpectations(t)
	})

	t.Run("token not exists", func(t *testing.T) {
		repo.On("GetUserData", uint(1)).Return(successGetUserData, nil).Once()
		repo.On("GetToken", successSend.Group_ID, successSend.User_ID).Return([]string{}).Once()
		data, err := usecase.SendNotification(successSend, client, ctx)
		assert.EqualError(t, err, "notification not sent")
		assert.Equal(t, 0, data)
		repo.AssertExpectations(t)
	})

	t.Run("token exists chats", func(t *testing.T) {
		repo.On("GetUserData", uint(1)).Return(successGetUserData, nil).Once()
		repo.On("GetToken", successSend.Group_ID, successSend.User_ID).Return([]string{"ePJbIm4SVq36C_lFMPGovY:APA91bGNnLsikuZH75zcLV5y-SUqYuvGmk6QNWy42BQNwGkqufiwKv9NULFqMczGCA-5jHAguloglkAZCFavypnjipSR2BpSKgW5bzB-tnBPucD2NPfurtfIjKw168QWdtHP0wNAChnF"}).Once()
		data, err := usecase.SendNotification(successSend, client, ctx)
		assert.Nil(t, err)
		assert.Equal(t, 1, data)
		repo.AssertExpectations(t)
	})

	successSend.IsSOS = true
	t.Run("token exists sos", func(t *testing.T) {
		repo.On("GetUserData", uint(1)).Return(successGetUserData, nil).Once()
		repo.On("GetToken", successSend.Group_ID, successSend.User_ID).Return([]string{"ePJbIm4SVq36C_lFMPGovY:APA91bGNnLsikuZH75zcLV5y-SUqYuvGmk6QNWy42BQNwGkqufiwKv9NULFqMczGCA-5jHAguloglkAZCFavypnjipSR2BpSKgW5bzB-tnBPucD2NPfurtfIjKw168QWdtHP0wNAChnF"}).Once()
		data, err := usecase.SendNotification(successSend, client, ctx)
		assert.Nil(t, err)
		assert.Equal(t, 1, data)
		repo.AssertExpectations(t)
	})
}
