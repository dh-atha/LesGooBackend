package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	e.POST("/group", handler.InsertGroup(), JWT)
	e.GET("/group/:id", handler.GetGroupByID())
	e.DELETE("/group/:id", handler.DeleteGroupByIDGroup(), JWT)
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

func (gh *groupHandler) InsertGroup() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := common.ExtractData(c)
		if id == -1 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		//	Genereate UUID for group_id
		group_id := uuid.New()

		tmp := Group{}

		errBind := c.Bind(&tmp)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		/*
			Code AWS S3 for upload image group
			tmp.GroupImg = "Ini file image group"
		*/

		tmp.ID = group_id.String()
		tmp.Created_By_User_ID = uint(id)

		err := gh.groupUsecase.AddGroup(ToModelGroup(tmp))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		tmp2 := GroupUser{}

		tmp2.GroupID = group_id.String()
		tmp2.UserID = uint(id)
		tmp2.Longitude = tmp.Longitude
		tmp2.Latitude = tmp.Latitude

		err2 := gh.groupUsecase.AddGroupUser(ToModelGroupUser(tmp2))
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":     201,
			"message":  "success operation",
			"id_group": tmp.ID,
		})

	}
}

func (gh *groupHandler) GetGroupByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		result, err := gh.groupUsecase.GetGroupDetail(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		response := FromModelByID(result)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success operation",
			"data":    response,
		})
	}
}

func (gh *groupHandler) DeleteGroupByIDGroup() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")

		id_user := common.ExtractData(c)
		if id_user == -1 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		err := gh.groupUsecase.DeleteGroupByID(id, uint(id_user))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success operation",
		})
	}
}
