/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 01:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 01:14:50
 */
package routes

import (
	"github.com/Anzz-bot/DouYin_demo/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetApiGroupRouters(router *gin.RouterGroup) {
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.GET("/testApi", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/register/", app.Register)
	router.POST("/login/", app.Login)

}
