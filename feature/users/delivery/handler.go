package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func New(e *echo.Echo, us domain.UserUsecase) {
	_ = &userHandler{
		userUsecase: us,
	}
	_ = middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
}
