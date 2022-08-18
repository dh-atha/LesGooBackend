package messaging

import (
	"context"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func InitFirebaseClient(ctx context.Context) *messaging.Client {
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
	return client
}

func SendChat(data domain.Chat, tokens []string, client *messaging.Client, context context.Context) (*messaging.BatchResponse, error) {
	message := &messaging.MulticastMessage{
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Title: data.Group_ID,
				Body:  data.Message,
			},
			Data: map[string]string{
				"action": "chat",
			},
			FcmOptions: &messaging.WebpushFcmOptions{
				Link: "https://google.com",
			},
		},
		Notification: &messaging.Notification{
			Title: data.Group_ID,
			Body:  data.Message,
		},
		Data: map[string]string{
			"action": "chat",
		},
		Tokens: tokens,
	}

	response, err := client.SendMulticast(context, message)
	return response, err
}

func SendLocation(tokens []string, client *messaging.Client, context context.Context) (*messaging.BatchResponse, error) {
	message := &messaging.MulticastMessage{
		Webpush: &messaging.WebpushConfig{
			Data: map[string]string{
				"action": "location",
			},
		},
		Data: map[string]string{
			"action": "location",
		},
		Tokens: tokens,
	}

	response, err := client.SendMulticast(context, message)
	if err != nil {
		log.Fatalln(err)
	}
	return response, err
}
