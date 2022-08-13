package usecase

import (
	"context"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
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

func (cd *chatData) SendNotification(data domain.Chat) (int, error) {
	res := cd.chatData.GetToken(data.Group_ID)

	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: config.ProjectID,
	}, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.MulticastMessage{
		// Webpush: &messaging.WebpushConfig{
		// 	Notification: &messaging.WebpushNotification{
		// 		Title: data.Group_ID,
		// 		Body:  data.Message,
		// 	},
		// },
		Notification: &messaging.Notification{
			Title: data.Group_ID,
			Body:  data.Message,
		},
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Tokens: res,
	}

	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	return response.SuccessCount, nil
}
