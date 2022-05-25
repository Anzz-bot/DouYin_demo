/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package models

type Relation struct {
	ID
	UserID   uint64 `json:"user_id"`   // 用户ID
	FollowID uint64 `json:"follow_id"` // 被关注用户ID
	Timestamps
	SoftDeletes
}
