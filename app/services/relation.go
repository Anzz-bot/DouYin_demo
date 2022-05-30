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

type relationService struct {
}

var RelationService = new(relationService)

func (relationService *relationService) Relation(params request.Relation) (err error) {
	switch params.ActionType {
	case 1:
		if err = relationService.RelationOn(params.ToUserId, params.UserId); err != nil {
			return errors.New("关注失败:" + err.Error())
		}
		return
	case 2:
		if err = relationService.RelationOff(params.ToUserId, params.UserId); err != nil {
			return errors.New("取关失败:" + err.Error())
		}
		return
	}
	return errors.New("获取参数失败")
}

func (relationService *relationService) RelationOn(toUserId uint64, userId uint64) (err error) {
	if toUserId == userId {
		return errors.New("不能关注自己哦")
	}

	if relationService.IsRelation(userId, toUserId) {
		return errors.New("已经关注过了，亲")
	}

	relation := &models.Relation{
		UserID:   userId,
		FollowID: toUserId,
	}

	// 在relation表中添加该记录
	if err = global.App.DB.Create(relation).Error; err != nil {
		return
	}
	// 更新user表中被关注用户的粉丝数
	if err = global.App.DB.Model(&models.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return
	}

	// 更新user表中发起关注用户的关注数
	if err = global.App.DB.Model(&models.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return
	}
	return
}

func (relationService *relationService) RelationOff(toUserId uint64, userId uint64) (err error) {
	if toUserId == userId {
		return errors.New("不能关注自己哦")
	}

	if !relationService.IsRelation(userId, toUserId) {
		return errors.New("还未关注过哦，亲")
	}
	// 在relation表中添加该记录
	if err = global.App.DB.Where("user_id=? AND follow_id=?", userId, toUserId).Delete(&models.Relation{}).Error; err != nil {
		return
	}
	// 更新user表中被关注用户的粉丝数
	if err = global.App.DB.Model(&models.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		return
	}

	// 更新user表中发起关注用户的关注数
	if err = global.App.DB.Model(&models.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		return
	}
	return
}

func (relationService *relationService) IsRelation(userId uint64, toUserId uint64) bool {
	return global.App.DB.First(&models.Relation{}, "user_id = ? and follow_id = ?", userId, toUserId).Error == nil
}

func (relationService *relationService) RelationFollowList(param request.RelationFollowList) (err error, relationFollowListResponse response.RelationFollowListResponse) {
	follows, err := relationService.GetFollowUserList(param.UserId, global.NowUserID)
	relationFollowListResponse = response.RelationFollowListResponse{
		UserList: follows,
	}
	if err != nil {
		return
	}
	return
}

func (relationService *relationService) RelationFollowerList(param request.RelationFollowerList) (err error, relationFollowerListResponse response.RelationFollowerListResponse) {
	follows, err := relationService.GetFollowerUserList(param.UserId, global.NowUserID)
	relationFollowerListResponse = response.RelationFollowerListResponse{
		UserList: follows,
	}
	if err != nil {
		return
	}
	return
}

func (relationService *relationService) GetFollowUserList(userID, requestFromID uint64) ([]*response.UserAPI, error) {
	var userList []*response.UserAPI
	if err := global.App.DB.Raw("SELECT\n    u.id AS id,\n    u.name AS name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow\nFROM users u\nLEFT JOIN relations r ON u.id = r.follow_id AND r.user_id=?\nWHERE u.id IN (\n    SELECT r2.follow_id\n    FROM relations r2\n    WHERE r2.user_id=?\n    )AND r.deleted_at IS NULL;", requestFromID, userID).Scan(&userList).Error; err != nil {
		global.App.Log.Info(err.Error())
		return nil, err
	}
	return userList, nil
}

func (relationService *relationService) GetFollowerUserList(userID, requestFromID uint64) ([]*response.UserAPI, error) {
	var userList []*response.UserAPI
	if err := global.App.DB.Raw("SELECT\n    u.id AS id,\n    u.name AS name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow\nFROM users u\nLEFT JOIN relations r ON u.id = r.follow_id AND r.user_id=?\nWHERE u.id IN (\n    SELECT r2.user_id\n    FROM relations r2\n    WHERE r2.follow_id=?\n    )AND r.deleted_at IS NULL;", requestFromID, userID).Scan(&userList).Error; err != nil {
		global.App.Log.Info(err.Error())
		return nil, err
	}
	return userList, nil
}
