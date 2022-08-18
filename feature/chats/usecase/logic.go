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
	userData, err := cd.chatData.GetUserData(data.User_ID)
	if err != nil {
		return 0, err
	}

	tokens := cd.chatData.GetToken(data.Group_ID, data.User_ID)
	log.Println(tokens)
	if len(tokens) < 1 {
		return 0, errors.New("notification not sent")
	}
	var response *messaging.BatchResponse
	if !data.IsSOS {
		response, _ = fcm.SendChat(data, tokens, client, ctx, userData)
	} else {
		response, _ = fcm.SendSOS(data, tokens, client, ctx, userData)
	}
	return response.SuccessCount, nil
}
