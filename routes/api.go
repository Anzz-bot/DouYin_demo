package routes

import (
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

}
