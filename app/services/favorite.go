package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"gorm.io/gorm"
)

type favoriteService struct {
}

var FavoriteService = new(favoriteService)

func (favoriteService *favoriteService) Favorite(params request.Favorite) (err error) {
	global.App.Log.Info(string(params.UserId))
	switch params.ActionType {
	case 1:
		if err = favoriteService.FavoriteOn(params.VideoId, params.UserId); err != nil {
			return errors.New("点赞失败:" + err.Error())
		}
		return
	case 2:
		if err = favoriteService.FavoriteOff(params.VideoId, params.UserId); err != nil {
			return errors.New("取消点赞失败" + err.Error())
		}
		return
	}
	return errors.New("获取参数失败")
}

func (favoriteService *favoriteService) FavoriteOn(videoId uint64, userId uint64) (err error) {
	favorite := &models.Favorite{
		UserID:  userId,
		VideoID: videoId,
	}
	global.App.Log.Info(string(userId))
	// 在favorite表中添加该记录
	if err = global.App.DB.Create(favorite).Error; err != nil {
		return
	}
	// 更新video表中的点赞数
	if err = global.App.DB.Model(&models.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
		return
	}
	return
}

func (favoriteService *favoriteService) FavoriteOff(videoId uint64, userId uint64) (err error) {

	// 在favorite表中添加该记录
	if err = global.App.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&models.Favorite{}).Error; err != nil {
		return
	}
	// 更新video表中的点赞数
	if err = global.App.DB.Model(&models.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
		return
	}
	return
}
