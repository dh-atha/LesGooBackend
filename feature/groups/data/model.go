package data

import (
	"lesgoobackend/domain"
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"
	usersdata "lesgoobackend/feature/users/data"
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Created_By_User_ID uint
	ID                 string `gorm:"type:VARCHAR(255);primaryKey"`
	Name               string
	Description        string
	Start_Date         string
	End_Date           string
	Start_Dest         string
	Final_Dest         string
	GroupImg           string
	Status             string
	Group_Users        []groupUsersData.Group_User `gorm:"foreignKey:Group_ID"`
	Chats              []chatsData.Chat            `gorm:"foreignKey:Group_ID"`
}

type Group_User struct {
	gorm.Model
	Group_ID  string
	User_ID   uint
	Longitude float64
	Latitude  float64
	User      usersdata.User `gorm:"foreignKey:User_ID"`
}

func fromModelGroup(model domain.Group) Group {
	return Group{
		Model:              gorm.Model{CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
		Created_By_User_ID: model.Created_By_User_ID,
		ID:                 model.ID,
		Name:               model.Name,
		Description:        model.Description,
		Start_Date:         model.Start_Date,
		End_Date:           model.End_Date,
		Start_Dest:         model.Start_Dest,
		Final_Dest:         model.Final_Dest,
		GroupImg:           model.GroupImg,
		Status:             model.Status,
	}
}

func fromModelGroupUser(model domain.Group_User) Group_User {
	return Group_User{
		Model: gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Group_ID:  model.Group_ID,
		User_ID:   model.User_ID,
		Longitude: model.Longitude,
		Latitude:  model.Latitude,
	}
}

func (data *Group) toDomainByID() domain.Group {
	return domain.Group{
		ID:          data.ID,
		Name:        data.Name,
		GroupImg:    data.GroupImg,
		Start_Date:  data.Start_Date,
		End_Date:    data.End_Date,
		Description: data.Description,
	}
}

func toDomainByID(data Group) domain.Group {
	return data.toDomainByID()
}

func (data *Group_User) toUsersDomain() domain.UsersbyID {
	return domain.UsersbyID{
		UserID:   data.User_ID,
		Username: data.User.Username,
	}
}

func ToUsersDomainList(data []Group_User) []domain.UsersbyID {
	result := []domain.UsersbyID{}
	for key := range data {
		result = append(result, data[key].toUsersDomain())
	}
	return result
}
