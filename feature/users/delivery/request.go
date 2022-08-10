package delivery

import "lesgoobackend/domain"

type InsertFormat struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (i InsertFormat) ToModel() domain.User {
	return domain.User{
		Username: i.Username,
		Email:    i.Email,
		Phone:    i.Phone,
		Password: i.Password,
	}
}
