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
