package delivery

import (
	"context"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"net/http"
	"strconv"

	"firebase.google.com/go/messaging"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type chatHandler struct {
	chatHandler domain.ChatUsecase
}

func New(e *echo.Echo, cu domain.ChatUsecase) {
	handler := &chatHandler{chatHandler: cu}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/chats", handler.SendChats(), useJWT)
}

func (ch *chatHandler) SendChats() echo.HandlerFunc {
	return func(c echo.Context) error {
		FCMClient := c.Get("FCM").(*messaging.Client)
		context := c.Get("context").(context.Context)

		var chatRequest SendChatRequest
		err := c.Bind(&chatRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		err = validator.New().Struct(chatRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		var data domain.Chat = chatRequest.ToDomain()
		data.User_ID = uint(common.ExtractData(c))

		err = ch.chatHandler.SendChats(data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		response, err := ch.chatHandler.SendNotification(data, FCMClient, context)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "successfully sent to: " + strconv.Itoa(response),
		})
	}
}
