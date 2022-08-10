package delivery

import "lesgoobackend/domain"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func FromModel(data domain.User) UserResponse {
	var res UserResponse
	res.ID = data.ID
	res.Username = data.Username
	res.Email = data.Email
	res.Phone = data.Phone
	res.Password = data.Password
	return res
}
