package response

type CommentResponse struct {
	Response
	CommentList []*Comment `json:"comment_list"`
}

type CommentListResponse struct {
	Response
	CommentList []*Comment `json:"comment_list"`
}

type Comment struct {
	ID         uint64   `json:"id"`                                        // 评论id
	User       *UserAPI `json:"user" gorm:"embedded;embeddedPrefix:user_"` // 评论用户
	Content    string   `json:"content"`                                   // 评论内容
	CreateDate string   `json:"create_date"`                               // 评论时间
}
