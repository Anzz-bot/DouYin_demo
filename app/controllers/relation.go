package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
	"github.com/gin-gonic/gin"
)

func Relation(c *gin.Context) {
	var form request.Relation
	user_id := global.NowUserID
	token := c.Query("token")
	to_user_id := utils.StrToUint64(c.Query("to_user_id"))
	action_type := utils.StrToInt32(c.Query("action_type"))
	form = request.Relation{
		UserId:     user_id,
		Token:      token,
		ToUserId:   to_user_id,
		ActionType: action_type,
	}
	if err := services.RelationService.Relation(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, response.ResponseSuccess)
	}

}
