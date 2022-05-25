/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Publish(c *gin.Context) {
	var form request.VideoUpload
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		response.BusinessFail(c, err.Error())
	}
	form = request.VideoUpload{
		Data:  data,
		Token: token,
		Title: title,
	}
	if err := services.PublishService.Publish(c, form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, response.ResponseSuccess)
	}

}

func PublishList(c *gin.Context) {
	var form request.PublishList
	token := c.Query("token")
	user_id := c.Query("user_id")
	form = request.PublishList{
		UserId: user_id,
		Token:  token,
	}
	if err, publishListResponse := services.PublishService.PublishList(form); err != nil {
		c.JSON(http.StatusBadRequest, publishListResponse)
	} else {
		c.JSON(http.StatusOK, publishListResponse)
	}

}
