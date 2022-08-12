package usecase

import (
	"context"
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
	opt := option.WithCredentialsFile("coba.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	// _ = "eKRB1IDFsdnJwbw-6PMIBD:APA91bHs6BPiLJ6UHNK7tii39xzFc3z9nLnlFus_fwglhsy6LrVIts9At9UrIR9MI7nOOiFlEHfxGD0eM2lE1eLwgLanEqScfjL7h19xv2rAj7rxiKQji1aFH01VAKsc60sgzyHKiepB"

	// See documentation on defining a message payload.
	message := &messaging.MulticastMessage{
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

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	return response.SuccessCount, nil
}
