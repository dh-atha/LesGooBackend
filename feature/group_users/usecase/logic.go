package usecase

import (
	"context"
	"errors"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type groupUsersData struct {
	groupUsersData domain.Group_UserData
}

func New(data domain.Group_UserData) domain.Group_UserUsecase {
	return &groupUsersData{
		groupUsersData: data,
	}
}

// // LeaveGroup implements domain.Group_UserUsecase
func (gud *groupUsersData) LeaveGroup(data domain.Group_User) error {
	if data.Group_ID == "" || data.User_ID == 0 {
		return errors.New("failed")
	}

	err := gud.groupUsersData.Leave(data)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

// AddJoined implements domain.Group_UserUsecase
func (gud *groupUsersData) AddJoined(data domain.Group_User) error {
	if data.Group_ID == "" || data.Latitude == 0 || data.Longitude == 0 {
		return errors.New("failed")
	}

	err := gud.groupUsersData.Joined(data)
	if err != nil {
		return errors.New("failed")
	}

	return nil
}

func (gud *groupUsersData) UpdateLocation(data domain.Group_User) error {
	err := gud.groupUsersData.Update(data)
	if err != nil {
		return err
	}

	res := gud.groupUsersData.GetToken(data.Group_ID, data.User_ID)
	log.Println(res)

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
		Data: map[string]string{
			"action": "update location",
		},
		Tokens: res,
	}

	response, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(response.Responses); i++ {
		log.Println(response.Responses[i])
	}

	return err
}
