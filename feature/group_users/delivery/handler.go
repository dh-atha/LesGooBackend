package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"net/http"

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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		tmp.UserID = uint(id)
		err := gu.groupUsersUsecase.AddJoined(ToModelJoin(tmp))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "success operation",
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		tmp.UserID = uint(id)
		err := gu.groupUsersUsecase.LeaveGroup(ToModeLeave(tmp))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "success operation",
		})

	}
}

func (gu *groupUsersHandler) UpdateLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		err = gu.groupUsersUsecase.UpdateLocation(ToModelJoin(req))
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
