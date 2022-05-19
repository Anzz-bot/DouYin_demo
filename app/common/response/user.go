package response

import (
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterResponse struct {
	Response
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

var ResponseSuccess = Response{
	0,
	"ok",
}

func RegisterSuccess(c *gin.Context, user models.User) {
	global.App.Log.Info(user.Token)

	c.JSON(http.StatusOK, RegisterResponse{
		ResponseSuccess,
		user.ID.ID,
		user.Token,
	})
}

func RegisterFail(c *gin.Context, errorCode int, msg string) {
	var ResponseFail = Response{
		errorCode,
		msg,
	}
	c.JSON(http.StatusOK, RegisterResponse{
		ResponseFail,
		0,
		"fail",
	})
}
