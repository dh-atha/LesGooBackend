package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type groupHandler struct {
	groupUsecase domain.GroupUsecase
}

func New(e *echo.Echo, gu domain.GroupUsecase) {
	_ = &groupHandler{
		groupUsecase: gu,
	}
	_ = middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
}
