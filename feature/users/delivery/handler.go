package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"

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
	JWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/register", handler.InsertUser())
	e.POST("/login", handler.LoginHandler())
	e.POST("/logout", handler.Logout(), JWT)
	e.PUT("/users", handler.UpdateUser(), JWT)
	e.GET("/users", handler.GetProfile(), JWT)
	e.DELETE("/users", handler.DeleteUser(), JWT)
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

		if tmp.Username == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Invalid Username",
			})
		} else if tmp.Email == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Invalid Email",
			})
		} else if tmp.Phone == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Invalid Phone",
			})
		} else if tmp.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Invalid Password",
			})
		}

		_, err := uh.userUsecase.AddUser(tmp.ToModel())
		if err != nil {
			if err.Error() == "Invalid Username" || err.Error() == "Invalid Email" || err.Error() == "Invalid Phone" {
				log.Println("Cannot proces data", err)
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    400,
					"message": "username or email or Telephone Number Already Exist",
				})
			} else {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": "Internal Server Error",
				})
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "success operation",
		})
	}
}

func (uh *userHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "cannot read input",
			})
		}
		err = validator.New().Struct(userLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": e.Error(),
			})
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "cannot read input",
			})
		}

		groupID := uh.userUsecase.GetGroupID(data)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success login",
			"data": map[string]interface{}{
				"token":     common.GenerateToken(int(data.ID)),
				"fcm_token": data.Fcm_Token,
				"group_id":  groupID,
			},
		})
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		session := c.Get("session").(*session.Session)
		bucket := c.Get("bucket").(string)
		var tmp UpdateFormat
		err := c.Bind(&tmp)
		id := common.ExtractData(c)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "wrong input",
			})
		}
		fileData, err := c.FormFile("profileimg")
		if fileData != nil {
			if err == http.ErrMissingFile || err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    400,
					"message": err.Error(),
				})
			}
			url, err := uh.userUsecase.UploadFiles(session, bucket, fileData)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    500,
					"message": err.Error(),
				})
			}
			tmp.ProfileImg = url
		} else {
			tmp.ProfileImg = ""
		}
		data, err := uh.userUsecase.UpdateUser(id, tmp.UpdateToModel())
		if err != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
				"data":    data,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success operation",
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"code":    400,
					"message": "data not found",
				})
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success Operation",
			"data":    data,
		})
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		_, errDel := uh.userUsecase.DeleteUser(id)
		if errDel != nil {
			if errDel.Error() == "leave group before deleting your account" {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    400,
					"message": errDel.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success Operation",
		})
	}
}

func (uh *userHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractData(c)
		err := uh.userUsecase.Logout(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success logout",
		})
	}
}
