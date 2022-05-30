package middleware

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		},
		response.ServerError)
}
