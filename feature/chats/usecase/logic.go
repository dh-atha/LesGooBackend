package usecase

import "lesgoobackend/domain"

type chatData struct {
	chatData domain.ChatData
}

func New(cd domain.ChatData) domain.ChatUsecase {
	return &chatData{
		chatData: cd,
	}
}
