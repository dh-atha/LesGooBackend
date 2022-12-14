package delivery

import (
	"lesgoobackend/config"
	"lesgoobackend/domain"
	"lesgoobackend/feature/common"
	"lesgoobackend/feature/middlewares"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
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
		session := c.Get("session").(*session.Session)
		bucket := c.Get("bucket").(string)

		id := common.ExtractData(c)
		if id == -1 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		//	Generate UUID for group_id
		group_id := uuid.New()

		tmp := Group{}

		errBind := c.Bind(&tmp)
		if errBind != nil {
			log.Println(errBind.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Bad Request",
			})
		}
		err := validator.New().Struct(tmp)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Bad Request",
			})
		}

		groupImg, errImg := c.FormFile("groupimg")
		if errImg != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": errImg.Error(),
			})
		}

		groupImgUrl, errImgUrl := gh.groupUsecase.UploadFiles(session, bucket, groupImg, group_id.String())
		if errImgUrl != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": errImgUrl.Error(),
			})
		}

		tmp.GroupImg = groupImgUrl

		tmp.ID = group_id.String()
		tmp.Created_By_User_ID = uint(id)

		err = gh.groupUsecase.AddGroup(ToModelGroup(tmp))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		tmp2 := GroupUser{}

		tmp2.GroupID = group_id.String()
		tmp2.UserID = uint(id)
		tmp2.Longitude = tmp.Longitude
		tmp2.Latitude = tmp.Latitude

		err2 := gh.groupUsecase.AddGroupUser(ToModelGroupUser(tmp2))
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":     http.StatusCreated,
			"group_id": tmp.ID,
			"message":  "Success Create New Groups",
		})

	}
}

func (gh *groupHandler) GetGroupByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		result, err := gh.groupUsecase.GetGroupDetail(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Invalid Id",
			})
		}

		response := FromModelByID(result)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success",
			"data":    response,
		})
	}
}

func (gh *groupHandler) DeleteGroupByIDGroup() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")

		id_user := common.ExtractData(c)
		if id_user == -1 {
			return c.JSON(http.StatusForbidden, "Access Forbidden")
		}

		err := gh.groupUsecase.DeleteGroupByID(id, uint(id_user))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success Operation",
		})
	}
}
