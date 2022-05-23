/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 20:14:50
 */
package response

import (
	"github.com/Anzz-bot/DouYin_demo/app/models"
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

func LoginSuccess(c *gin.Context, user models.User) {
	
	c.JSON(http.StatusOK, RegisterResponse{
		ResponseSuccess,
		user.ID.ID,
		user.Token,
	})
}

func LoginFail(c *gin.Context, errorCode int, msg string) {
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