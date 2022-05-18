/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 17:11:15
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 17:11:15
 */

package global

import (
	"github.com/Anzz-bot/DouYin_demo/config"
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
}

var App = new(Application)
