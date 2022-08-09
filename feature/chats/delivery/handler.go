package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type chatHandler struct {
	chatHandler domain.ChatUsecase
}

func New(e *echo.Echo, cu domain.ChatUsecase) {
	_ = &chatHandler{chatHandler: cu}
	_ = middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
}
