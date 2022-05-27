/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Favorite(c *gin.Context) {
	var form request.Favorite
	user_id := global.NowUserID

	token := c.Query("token")
	video_id := utils.StrToUint64(c.Query("video_id"))
	action_type := utils.StrToInt32(c.Query("action_type"))
	form = request.Favorite{
		UserId:     user_id,
		Token:      token,
		VideoId:    video_id,
		ActionType: action_type,
	}
	if err := services.FavoriteService.Favorite(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, response.ResponseSuccess)
	}

}

func FavoriteList(c *gin.Context) {
	var form request.FavoriteList
	token := c.Query("token")
	user_id := c.Query("user_id")
	form = request.FavoriteList{
		UserId: utils.StrToUint64(user_id),
		Token:  token,
	}
	if err, favoriteListResponse := services.FavoriteService.FavoriteList(form); err != nil {
		c.JSON(http.StatusBadRequest, favoriteListResponse)
	} else {
		c.JSON(http.StatusOK, favoriteListResponse)
	}

}
