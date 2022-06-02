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
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/jassue/go-storage/storage"
	uuid "github.com/satori/go.uuid"
	"path"
)

type publishService struct {
}

var PublishService = new(publishService)

const videoCacheKeyPre = "video:"

// 文件存储目录
func (publishService *publishService) makeFaceDir(title string) string {
	return global.App.Config.App.Env + "/" + title
}

// HashName 生成文件名称（使用 uuid）
func (publishService *publishService) HashName(fileName string) string {
	fileSuffix := path.Ext(fileName)
	return uuid.NewV4().String() + fileSuffix
}

// Publish 视频投稿
func (publishService *publishService) Publish(c *gin.Context, params request.VideoUpload) (err error) {
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
	key := publishService.makeFaceDir(params.Title) + "/" + publishService.HashName(params.Data.Filename)
	disk := global.App.Disk()
	err = disk.Put(localPrefix+key, file, params.Data.Size)
	if err != nil {
		return
	}
	global.App.Log.Info(localPrefix + key)
	utils.GetSnapshot("./storage/app/"+localPrefix+key, "./storage/app/"+localPrefix+"local/"+params.Title+"/cover", 1)
	video := models.Video{
		AuthorID: global.NowUserID,
		DiskType: string(global.App.Config.Storage.Default),
		SrcType:  1,
		PlayUrl:  global.App.Config.App.AppUrl + ":" + global.App.Config.App.Port + "/storage/" + key, //"http://192.168.0.107:8888/storage/"
		CoverUrl: global.App.Config.App.AppUrl + ":" + global.App.Config.App.Port + "/storage/local/" + params.Title + "/cover.jpeg",
	}
	err = global.App.DB.Create(&video).Error
	if err != nil {
		return
	}

	return
}

func (publishService *publishService) PublishList(param request.PublishList) (err error, publishListResponse response.PublishListResponse) {
	videoAuthorApis, err := publishService.GetPublishListByAuthorID(param.UserId)
	publishListResponse = response.PublishListResponse{
		VideoList: videoAuthorApis,
	}
	if err != nil {
		return err, publishListResponse
	}
	return nil, publishListResponse
}

func (publishService *publishService) GetPublishListByAuthorID(authorID uint64) ([]*response.VideoAuthorApi, error) {
	userID := authorID // 这里的UserID是发起请求的用户ID
	var VideoAuthorApis []*response.VideoAuthorApi
	// 下述查询预留了发起请求的用户ID与视频作者ID不一致的情况
	// 如果发起请求的用户ID与视频作者ID不一致，需要传入userID
	err := global.App.DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    IF(r.id IS NULL,false,true) AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n     IF(f.id IS NULL,false,true) AS is_favorite\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nLEFT JOIN favorites f on u.id = f.user_id AND v.id = f.video_id\nWHERE author_id = ?\nORDER BY v.created_at", userID, authorID).Scan(&VideoAuthorApis).Error
	if err != nil {
		return nil, err
	}
	return VideoAuthorApis, nil
}
