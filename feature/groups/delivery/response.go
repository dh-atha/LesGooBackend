package delivery

import "lesgoobackend/domain"

type GroupByID struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	GroupImg    string      `json:"groupimg"`
	Start_Date  string      `json:"start_date"`
	End_Date    string      `json:"end_date"`
	Description string      `json:"description"`
	UsersbyID   []UsersbyID `json:"group_users"`
}

type UsersbyID struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

func FromModelByID(data domain.Group) GroupByID {
	return GroupByID{
		ID:          data.ID,
		Name:        data.Name,
		GroupImg:    data.GroupImg,
		Start_Date:  data.Start_Date,
		End_Date:    data.End_Date,
		Description: data.Description,
		UsersbyID:   FromGroupUsersModelList(data.UsersbyID),
	}
}

func FromGroupUsersModel(data domain.UsersbyID) UsersbyID {
	return UsersbyID{
		UserID:   data.UserID,
		Username: data.Username,
	}
}

func FromGroupUsersModelList(data []domain.UsersbyID) []UsersbyID {
	result := []UsersbyID{}
	for key := range data {
		result = append(result, FromGroupUsersModel(data[key]))
	}
	return result
}
