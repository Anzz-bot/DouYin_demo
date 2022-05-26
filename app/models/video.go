/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 02:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 02:53:50
 */
package models

type Video struct {
	ID
	AuthorID      uint64 `json:"author_id" gorm:"not null;comment:作者id"`
	PlayUrl       string `json:"play_url"  gorm:"not null;default:'';comment:视频地址"`
	CoverUrl      string `json:"cover_url" gorm:"not null;default:'';comment:封面地址"`
	FavoriteCount uint64 `json:"favorite_count" gorm:"not null;default:0;comment:点赞总数"`
	CommentCount  uint64 `json:"comment_count" gorm:"not null;default:0;comment:评论总数"`
	IsFavorite    bool   `json:"is_favorite" gorm:"not null;default:false;comment:是否点赞"`
	DiskType      string `json:"disk_type" gorm:"size:20;index;not null;comment:存储类型"`
	SrcType       int8   `json:"src_type" gorm:"not null;comment:链接类型 1相对路径 2外链"`
	Timestamps
	SoftDeletes
}
