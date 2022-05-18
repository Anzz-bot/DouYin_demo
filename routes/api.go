/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 01:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 01:14:50
 */
package routes

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
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
	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

}
