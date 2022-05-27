/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package request

type Comment struct {
	UserId      uint64 `form:"user_id" json:"user_id" `
	VideoId     uint64 `form:"video_id" json:"video_id" `
	CommentId   uint64 `form:"comment_id" json:"comment_id" `
	CommentText string `form:"comment_text" json:"comment_text" `
	Token       string `form:"token" json:"token" `
	ActionType  int32  `form:"action_type" json:"action_type" `
}

type CommentList struct {
	VideoId uint64 `form:"video_id" json:"video_id" `
	Token   string `form:"token" json:"token" `
}
