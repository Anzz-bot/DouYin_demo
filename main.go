/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 16:22:29
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 02:11:12
 */
 
package main

import (
	"github.com/Anzz-bot/DouYin_demo/bootstrap"
	"github.com/Anzz-bot/DouYin_demo/global"
)

func main() {

	bootstrap.InitializeConfig()    //load config
	bootstrap.InitializeValidator() //load validator
	//load log
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	global.App.Redis = bootstrap.InitializeRedis()
	global.App.Log.Info("redis init success!")
	//load DB
	global.App.DB = bootstrap.InitializeDB()
	global.App.Log.Info("DB init success!")

	//load Storage
	bootstrap.InitializeStorage()

	//close DB
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	bootstrap.RunServer()
}
