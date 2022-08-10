package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type groupUsersHandler struct {
	groupUsersUsecase domain.Group_UserUsecase
}

func New(e *echo.Echo, gus domain.Group_UserUsecase) {
	_ = &groupUsersHandler{
		groupUsersUsecase: gus,
	}
	_ = middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
}
