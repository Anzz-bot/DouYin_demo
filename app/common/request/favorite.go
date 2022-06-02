/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package request

type Favorite struct {
	UserId     uint64 `form:"user_id" json:"user_id" `
	VideoId    uint64 `form:"video_id" json:"video_id" `
	Token      string `form:"token" json:"token" `
	ActionType int32  `form:"action_type" json:"action_type" `
}

type FavoriteList struct {
	UserId uint64 `form:"user_id" json:"user_id" `
	Token  string `form:"token" json:"token" `
}
