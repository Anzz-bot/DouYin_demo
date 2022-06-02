/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package response

// RelationFollowListResponse 获取关注列表响应
type RelationFollowListResponse struct {
	Response
	UserList []*UserAPI `json:"user_list"` // 用户信息列表
}

// RelationFollowerListResponse 获取粉丝列表响应
type RelationFollowerListResponse struct {
	Response
	UserList []*UserAPI `json:"user_list"` // 用户信息列表
}
