/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"gorm.io/gorm"
)

type favoriteService struct {
}

var FavoriteService = new(favoriteService)

func (favoriteService *favoriteService) Favorite(params request.Favorite) (err error) {

	switch params.ActionType {
	case 1:
		if err = favoriteService.FavoriteOn(params.VideoId, params.UserId); err != nil {
			return errors.New("点赞失败:" + err.Error())
		}
		return
	case 2:
		if err = favoriteService.FavoriteOff(params.VideoId, params.UserId); err != nil {
			return errors.New("取赞失败:" + err.Error())
		}
		return
	}
	return errors.New("获取参数失败")
}

func (favoriteService *favoriteService) FavoriteOn(videoId uint64, userId uint64) (err error) {
	// 重复点赞
	if favoriteService.IsFavorite(userId, videoId) {
		return errors.New("已经点赞过了，亲")
	}

	favorite := &models.Favorite{
		UserID:  userId,
		VideoID: videoId,
	}

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

	if !favoriteService.IsFavorite(userId, videoId) {
		return errors.New("还未点赞过，亲")
	}
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

func (favoriteService *favoriteService) IsFavorite(userId uint64, videoId uint64) bool {
	return global.App.DB.First(&models.Favorite{}, "user_id = ? and video_id = ?", userId, videoId).Error == nil
}

func (favoriteService *favoriteService) FavoriteList(params request.FavoriteList) (err error, favoriteListResponse response.FavoriteListResponse) {
	var favorites []*response.VideoAuthorApi
	err = global.App.DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    IF(r.id IS NULL,false,true) AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    true AS is_favorite\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nWHERE v.id IN(\n    SELECT video_id\n    FROM favorites\n    WHERE user_id = ?\n    )\nORDER BY v.created_at;", params.UserId, params.UserId).Scan(&favorites).Error
	favoriteListResponse = response.FavoriteListResponse{
		VideoList: favorites,
	}
	if err != nil {
		return
	}
	return
}
