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
