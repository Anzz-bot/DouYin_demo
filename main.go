/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:22:29
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 16:22:29
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//test gin router
	//success return OK
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	//start gin server
	r.Run(":8080")
}
