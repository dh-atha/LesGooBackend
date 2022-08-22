package data

import (
	"lesgoobackend/domain"
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"

	// groupData "lesgoobackend/feature/groups/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Email      string
	Password   string
	Phone      string
	ProfileImg string `gorm:"default:https://lesgooproject.s3.ap-southeast-1.amazonaws.com/profileimg/user.png"`
	Fcm_Token  string
	// Groups      []groupData.Group           `gorm:"foreignKey:Created_By_User_ID"`	(Ini kalau tidak dibutuhkan, dihapus aja, soalnya import cycle dari aku (faza))
	Group_Users []groupUsersData.Group_User `gorm:"foreignKey:User_ID"`
	Chats       []chatsData.Chat            `gorm:"foreignKey:User_ID"`
}

func FromModel(data domain.User) User {
	var res User
	res.ProfileImg = data.ProfileImg
	res.Username = data.Username
	res.Password = data.Password
	res.Email = data.Email
	res.Phone = data.Phone
	res.Fcm_Token = data.Fcm_Token
	return res
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:         u.ID,
		ProfileImg: u.ProfileImg,
		Username:   u.Username,
		Password:   u.Password,
		Email:      u.Email,
		Phone:      u.Phone,
		Fcm_Token:  u.Fcm_Token,
	}
}
