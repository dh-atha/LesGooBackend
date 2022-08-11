package factory

import (
	chatData "lesgoobackend/feature/chats/data"
	chatDelivery "lesgoobackend/feature/chats/delivery"
	chatUsecase "lesgoobackend/feature/chats/usecase"
	group_userData "lesgoobackend/feature/group_users/data"
	group_userDelivery "lesgoobackend/feature/group_users/delivery"
	group_userUsecase "lesgoobackend/feature/group_users/usecase"
	groupData "lesgoobackend/feature/groups/data"
	groupDelivery "lesgoobackend/feature/groups/delivery"
	groupUsecase "lesgoobackend/feature/groups/usecase"
	userData "lesgoobackend/feature/users/data"
	userDelivery "lesgoobackend/feature/users/delivery"
	userUsecase "lesgoobackend/feature/users/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	userData := userData.New(db)
	validator := validator.New()
	userUsecase := userUsecase.New(userData, validator)
	userDelivery.New(e, userUsecase)

	groupData := groupData.New(db)
	groupUsecase := groupUsecase.New(groupData)
	groupDelivery.New(e, groupUsecase)

	group_userData := group_userData.New(db)
	group_userUsecase := group_userUsecase.New(group_userData)
	group_userDelivery.New(e, group_userUsecase)

	chatData := chatData.New(db)
	chatUsecase := chatUsecase.New(chatData)
	chatDelivery.New(e, chatUsecase)
}
