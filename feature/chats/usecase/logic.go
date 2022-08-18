package usecase

import (
	"context"
	"errors"
	"lesgoobackend/domain"
	"log"

	fcm "lesgoobackend/infrastructure/firebase/messaging"

	"firebase.google.com/go/messaging"
)

type chatData struct {
	chatData domain.ChatData
}

func New(cd domain.ChatData) domain.ChatUsecase {
	return &chatData{
		chatData: cd,
	}
}

func (cd *chatData) SendChats(data domain.Chat) error {
	err := cd.chatData.Insert(data)
	return err
}

func (cd *chatData) SendNotification(data domain.Chat, client *messaging.Client, ctx context.Context) (int, error) {
	tokens := cd.chatData.GetToken(data.Group_ID, data.User_ID)
	log.Println(tokens)
	if len(tokens) < 1 {
		return 0, errors.New("notification not sent")
	}
	response, _ := fcm.SendChat(data, tokens, client, ctx)
	return response.SuccessCount, nil
}
