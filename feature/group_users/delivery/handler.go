package delivery

import (
	"context"
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type groupUsersHandler struct {
	groupUsersUsecase domain.Group_UserUsecase
}

func New(e *echo.Echo, gus domain.Group_UserUsecase) {
	handler := &groupUsersHandler{
		groupUsersUsecase: gus,
	}
	JWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/group/join", handler.UserJoined(), JWT)
	e.POST("/group/leave", handler.LeaveGroup(), JWT)
	e.POST("/locations", handler.UpdateLocation(), JWT)
}

func (gu *groupUsersHandler) UserJoined() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		if id == -1 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		tmp := GroupUsers{}
		errBind := c.Bind(&tmp)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": errBind.Error(),
			})
		}

		err := validator.New().Struct(tmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		tmp.UserID = uint(id)
		err = gu.groupUsersUsecase.AddJoined(ToModelJoin(tmp))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success Join Group",
		})

	}
}

func (gu *groupUsersHandler) LeaveGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		if id == -1 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		tmp := GroupUsers{}

		errBind := c.Bind(&tmp)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
		}

		tmp.UserID = uint(id)
		err := gu.groupUsersUsecase.LeaveGroup(ToModeLeave(tmp))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "ok",
		})

	}
}

func (gu *groupUsersHandler) UpdateLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		FCMClient := c.Get("FCM").(*messaging.Client)
		context := c.Get("context").(context.Context)

		var req GroupUsers
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		err = validator.New().Struct(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		req.UserID = uint(common.ExtractData(c))

		err = gu.groupUsersUsecase.UpdateLocation(ToModelJoin(req), FCMClient, context)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success update location",
		})
	}
}
