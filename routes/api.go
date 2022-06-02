/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 01:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 01:14:50
 */
package routes

import (
	"github.com/Anzz-bot/DouYin_demo/app/controllers"

	"github.com/Anzz-bot/DouYin_demo/app/middleware"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetUserApiGroupRouters(router *gin.RouterGroup) {
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.GET("/testApi", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/register/", app.Register)
	router.POST("/login/", app.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/", app.Info)
	}

}

func SetPublishApiGroupRouters(router *gin.RouterGroup) {
	//authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	//{
	router.POST("/action/", app.Publish)
	router.GET("/list/", app.PublishList)
	//}
}

func SetFavoriteApiGroupRouters(router *gin.RouterGroup) {

	router.POST("/action/", app.Favorite)
	router.GET("/list/", app.FavoriteList)

}

func SetCommentApiGroupRouters(router *gin.RouterGroup) {

	router.POST("/action/", app.Comment)
	router.GET("/list/", app.CommentList)

}

func SetRelationApiGroupRouters(router *gin.RouterGroup) {

	router.POST("/action/", app.Relation)
	router.GET("/follow/list/", app.RelationFollowList)
	router.GET("/follower/list/", app.RelationFollowerList)
}
