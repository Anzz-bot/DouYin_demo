/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package models

type Comment struct {
	ID
	UserID      uint64 `json:"user_id"`      // 用户ID
	VideoID     uint64 `json:"video_id"`     // 视频ID
	CommentText string `json:"comment_text"` // 评论内容
	Timestamps
	SoftDeletes
}
