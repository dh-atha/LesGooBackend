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

type LoginFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		Username: lf.Username,
		Password: lf.Password,
	}
}

type UpdateFormat struct {
	ProfileImg string `json:"profileimg" form:"profileimg"`
	Username   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
}

func (uf *UpdateFormat) UpdateToModel() domain.User {
	return domain.User{
		ProfileImg: uf.ProfileImg,
		Username:   uf.Username,
		Email:      uf.Email,
		Phone:      uf.Phone,
	}
}
