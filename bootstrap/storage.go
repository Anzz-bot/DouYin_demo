/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package bootstrap

import (
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/jassue/go-storage/kodo"
	"github.com/jassue/go-storage/local"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	//	_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
