package delivery

import "lesgoobackend/domain"

type SendChatRequest struct {
	Group_ID string `json:"group_id" form:"group_id" validate:"required"`
	Message  string `json:"message" form:"message" validate:"required"`
	IsSOS    bool   `json:"isSOS" form:"isSOS"`
}

func (sc *SendChatRequest) ToDomain() domain.Chat {
	return domain.Chat{
		Group_ID: sc.Group_ID,
		Message:  sc.Message,
		IsSOS:    sc.IsSOS,
	}
}
