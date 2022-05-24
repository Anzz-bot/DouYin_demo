package models

type Video struct {
	ID
	AuthorID      uint64 `json:"author_id" gorm:"not null;comment:作者id"`
	PlayUrl       string `json:"play_url"  gorm:"not null;default:'';comment:视频地址"`
	CoverUrl      string `json:"cover_url" gorm:"not null;default:'';comment:封面地址"`
	FavoriteCount uint64 `json:"favorite_count" gorm:"not null;default:0;comment:点赞总数"`
	CommentCount  uint64 `json:"comment_count" gorm:"not null;default:0;comment:评论总数"`
	IsFavorite    bool   `json:"is_favorite" gorm:"not null;default:false;comment:是否点赞"`
	Timestamps
	SoftDeletes
}
