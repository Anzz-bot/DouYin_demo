/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 03:10:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 03:10:11
 */
package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var form request.Register

	username := c.Query("username")
	password := c.Query("password")

	form = request.Register{
		Name:     username,
		Password: password,
	}
	services.UserService.Register(form)

}
