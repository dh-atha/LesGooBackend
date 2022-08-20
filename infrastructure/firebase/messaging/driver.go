package messaging

import (
	"context"
	"fmt"
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

func SendChat(data domain.Chat, tokens []string, client *messaging.Client, context context.Context, userData domain.User) (*messaging.BatchResponse, error) {
	message := &messaging.MulticastMessage{
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Title: userData.Username,
				Body:  data.Message,
				Icon:  userData.ProfileImg,
				Image: "https://lesgooproject.s3.ap-southeast-1.amazonaws.com/logo.png",
			},
			Data: map[string]string{
				"action":     "chat",
				"username":   userData.Username,
				"message":    data.Message,
				"profileimg": userData.ProfileImg,
				"isSOS":      "false",
			},
			FcmOptions: &messaging.WebpushFcmOptions{
				Link: "https://google.com",
			},
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

func SendSOS(data domain.Chat, tokens []string, client *messaging.Client, context context.Context, userData domain.User) (*messaging.BatchResponse, error) {
	message := &messaging.MulticastMessage{
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Title:              fmt.Sprint("SOS by: ", userData.Username),
				Body:               data.Message,
				Icon:               userData.ProfileImg,
				Image:              "https://lesgooproject.s3.ap-southeast-1.amazonaws.com/logo.png",
				RequireInteraction: true,
			},
			Data: map[string]string{
				"action":     "sos",
				"username":   userData.Username,
				"message":    data.Message,
				"profileimg": userData.ProfileImg,
				"isSOS":      "true",
			},
			FcmOptions: &messaging.WebpushFcmOptions{
				Link: "https://google.com",
			},
		},
		Tokens: tokens,
	}

	response, err := client.SendMulticast(context, message)
	return response, err
}
