package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type groupHandler struct {
	groupUsecase domain.GroupUsecase
}

func New(e *echo.Echo, gu domain.GroupUsecase) {
	handler := &groupHandler{
		groupUsecase: gu,
	}
	JWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/group/chats", handler.GetChatsAndUsersLocation(), JWT)
}

func (gh *groupHandler) GetChatsAndUsersLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GetChatsAndUsersLocationRequest
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

		res, err := gh.groupUsecase.GetChatsAndUsersLocation(req.Group_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get chats",
			"data":    res,
		})
	}
}
