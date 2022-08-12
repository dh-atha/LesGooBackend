package delivery

type GetChatsAndUsersLocationRequest struct {
	Group_ID string `json:"group_id" form:"group_id" validate:"required"`
}
