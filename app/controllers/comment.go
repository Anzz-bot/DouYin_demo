package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Comment(c *gin.Context) {
	var form request.Comment
	user_id := global.NowUserID

	token := c.Query("token")
	video_id := utils.StrToUint64(c.Query("video_id"))
	comment_id := utils.StrToUint64(c.Query("comment_id"))
	action_type := utils.StrToInt32(c.Query("action_type"))
	comment_text := c.Query("comment_text")
	form = request.Comment{
		UserId:      user_id,
		Token:       token,
		VideoId:     video_id,
		ActionType:  action_type,
		CommentId:   comment_id,
		CommentText: comment_text,
	}
	if err, commentResponse := services.CommentService.Comment(form); err != nil {
		c.JSON(http.StatusBadRequest, commentResponse)
	} else {
		c.JSON(http.StatusOK, commentResponse)
	}

}

func CommentList(c *gin.Context) {
	var form request.CommentList
	token := c.Query("token")
	videoId := c.Query("video_id")
	form = request.CommentList{
		VideoId: utils.StrToUint64(videoId),
		Token:   token,
	}
	if err, commentListResponse := services.CommentService.CommentList(form); err != nil {
		c.JSON(http.StatusBadRequest, commentListResponse)
	} else {
		c.JSON(http.StatusOK, commentListResponse)
	}

}
