/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 17:11:15
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 17:11:15
 */


package global

import (
	"github.com/Anzz-bot/DouYin_demo/config"
	"github.com/go-redis/redis/v8"
	"github.com/jassue/go-storage/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//Store variables when the project is started
type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)

var NowUserID uint64

func (app *Application) Disk(disk ...string) storage.Storage {
	// 若未传参，默认使用配置文件驱动
	diskName := app.Config.Storage.Default
	if len(disk) > 0 {
		diskName = storage.DiskName(disk[0])
	}
	s, err := storage.Disk(diskName)
	if err != nil {
		panic(err)
	}
	return s
}
