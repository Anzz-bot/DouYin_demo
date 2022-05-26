/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package response

import (
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterLoginResponse struct {
	Response
	UserId uint64 `json:"user_id"`
	Token  string `json:"token"`
}
type InfoResponse struct {
	Response
	User UserAPI `json:"user"`
}
type UserAPI struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`           // 用户名称
	FollowCount   uint64 `json:"follow_count"`   // 关注总数
	FollowerCount uint64 `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // 是否关注
}

func RegisterSuccess(c *gin.Context, user models.User) {

	c.JSON(http.StatusOK, RegisterLoginResponse{
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
	c.JSON(http.StatusOK, RegisterLoginResponse{
		ResponseFail,
		0,
		"fail",
	})
}

func LoginSuccess(c *gin.Context, user models.User) {

	c.JSON(http.StatusOK, RegisterLoginResponse{
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
	c.JSON(http.StatusOK, RegisterLoginResponse{
		ResponseFail,
		0,
		"fail",
	})
}

func InfoSuccess(c *gin.Context, user models.User) {

	c.JSON(http.StatusOK, InfoResponse{
		ResponseSuccess,
		UserAPI{
			user.ID.ID,
			user.Name,
			user.FollowCount,
			user.FollowerCount,
			user.IsFollow,
		},
	})
}

func InfoFail(c *gin.Context, errorCode int, msg string) {
	var ResponseFail = Response{
		errorCode,
		msg,
	}
	c.JSON(http.StatusOK, RegisterLoginResponse{
		ResponseFail,
		0,
		"fail",
	})
}
