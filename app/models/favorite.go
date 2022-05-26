/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package models

type Favorite struct {
	ID
	UserID  uint64 `json:"user_id"`  // 点赞的用户ID
	VideoID uint64 `json:"video_id"` // 被点赞的视频ID
	Timestamps
	SoftDeletes
}
