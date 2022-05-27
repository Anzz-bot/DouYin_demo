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

type commentService struct {
}

var CommentService = new(commentService)

func (commentService *commentService) Comment(params request.Comment) (err error, commentResponse response.CommentResponse) {
	comments, err := commentService.GetCommentUserBundle(global.NowUserID, params.VideoId)
	commentResponse = response.CommentResponse{
		CommentList: comments,
	}
	switch params.ActionType {
	case 1:
		if err = commentService.CommentOn(params.UserId, params.VideoId, params.CommentText); err != nil {
			err = errors.New("评论失败：" + err.Error())
			return
		}
		return
	case 2:
		if err = commentService.CommentOff(params.UserId, params.VideoId, params.CommentId); err != nil {
			err = errors.New("删除评论失败:" + err.Error())
			return
		}
		return
	}
	err = errors.New("获取参数失败")
	return
}

func (commentService *commentService) CommentOn(userId uint64, videoId uint64, text string) (err error) {

	comment := &models.Comment{
		UserID:      userId,
		VideoID:     videoId,
		CommentText: text,
	}

	// 在comment表中添加该记录
	if err = global.App.DB.Create(comment).Error; err != nil {
		return
	}
	//
	if err = global.App.DB.Model(&models.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
		return
	}
	return
}

func (commentService *commentService) CommentOff(userId uint64, videoId uint64, CommentId uint64) (err error) {

	// 在favorite表中添加该记录
	if err = global.App.DB.Delete(&models.Comment{}, "user_id = ? and video_id = ? and id = ?", userId, videoId, CommentId).Error; err != nil {
		return
	}
	// 更新video表中的点赞数
	if err = global.App.DB.Model(&models.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
		return
	}
	return
}

// GetCommentList 获取评论列表
func (commentService *commentService) CommentList(params request.CommentList) (err error, commentListResponse response.CommentListResponse) {
	comments, err := commentService.GetCommentUserBundle(global.NowUserID, params.VideoId)
	commentListResponse = response.CommentListResponse{
		CommentList: comments,
	}
	if err != nil {
		global.App.Log.Info(err.Error())
		return
	}
	return
}

func (commentService *commentService) GetCommentUserBundle(userId uint64, videoId uint64) ([]*response.Comment, error) {
	var comments []*response.Comment
	err := global.App.DB.Raw("SELECT\n    c.id AS id,\n    u.id AS user_id,\n    u.name AS user_name,\n    u.follow_count AS follow_count,\n    u.follower_count AS follower_count,\n    IF(r.id IS NULL,false,true) AS is_follow,\n    c.comment_text AS content,\n    c.created_at AS create_date\nFROM comments c\nLEFT JOIN users u ON c.user_id=u.id\nLEFT JOIN relations r on u.id = r.follow_id AND r.user_id = ?\nWHERE c.video_id=?\nORDER BY c.updated_at;", userId, videoId).Scan(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
