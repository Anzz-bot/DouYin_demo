/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
	"github.com/jassue/go-storage/storage"
	uuid "github.com/satori/go.uuid"
	"path"
)

type publishSearvice struct {
}

var PublishService = new(publishSearvice)

const videoCacheKeyPre = "video:"

// 文件存储目录
func (publishSearvice *publishSearvice) makeFaceDir(title string) string {
	return global.App.Config.App.Env + "/" + title
}

// HashName 生成文件名称（使用 uuid）
func (publishSearvice *publishSearvice) HashName(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return uuid.NewV4().String() + fileSuffix
}

// Publish 视频投稿
func (publishSearvice *publishSearvice) Publish(c *gin.Context, params request.VideoUpload) (err error) {
	file, err := params.Data.Open()
	defer file.Close()
	if err != nil {
		err = errors.New("上传视频失败")
		return
	}

	localPrefix := ""
	// 本地文件存放路径为 storage/app/public
	// 配置了静态资源处理路由 router.Static("/storage", "./storage/app/public")
	// 所以此处不需要将 public/ 存入到 mysql 中，防止后续拼接文件 Url 错误
	if global.App.Config.Storage.Default == storage.Local {
		localPrefix = "public" + "/"
	}
	key := publishSearvice.makeFaceDir(params.Title) + "/" + publishSearvice.HashName(params.Data.Filename)
	disk := global.App.Disk()
	err = disk.Put(localPrefix+key, file, params.Data.Size)
	if err != nil {
		return
	}

	video := models.Video{
		AuthorID: global.NowUserID,
		DiskType: string(global.App.Config.Storage.Default),
		SrcType:  1,
		PlayUrl:  global.App.Config.App.AppUrl + ":" + global.App.Config.App.Port + "/storage/" + key, //"http://192.168.0.107:8888/storage/"
	}
	err = global.App.DB.Create(&video).Error
	if err != nil {
		return
	}

	return
}
