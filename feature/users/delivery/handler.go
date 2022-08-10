package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/middlewares"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func New(e *echo.Echo, us domain.UserUsecase) {
	handler := &userHandler{
		userUsecase: us,
	}
	_ = middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/register", handler.InsertUser())
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		errBind := c.Bind(&tmp)

		if errBind != nil {
			log.Println("Cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		_, err := uh.userUsecase.AddUser(tmp.ToModel())
		if err != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "success operation",
		})
	}
}
