/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 02:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 02:53:50
 */
package app


import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"net/http"

	"github.com/Anzz-bot/DouYin_demo/utils"
	"github.com/gin-gonic/gin"
	"time"
)


func Feed(c *gin.Context) {
	var form request.Feed
	latestTime := utils.StrToInt64(c.Query("latest_time"))
	if latestTime == 0 {
		latestTime = time.Now().UnixMilli()
	}
	token := c.Query("token")
	form = request.Feed{
		LatestTime: latestTime,
		Token:      token,
	}
	if err, feedresponse := services.FeedService.Feed(form); err != nil {
		c.JSON(http.StatusBadRequest, feedresponse)
	} else {
		c.JSON(http.StatusOK, feedresponse)
	}

}
