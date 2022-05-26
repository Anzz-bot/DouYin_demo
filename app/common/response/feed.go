/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 02:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 02:53:50
 */
package response

type FeedResponse struct {
	Response
	NextTime  int64             `json:"next_time"`
	VideoList []*VideoAuthorApi `json:"video_list"`
}

type VideoAuthorApi struct {
	ID            uint64   `json:"id"`
	Author        *UserAPI `json:"author" gorm:"embedded;embeddedPrefix:author_"`
	PlayURL       string   `json:"play_url"`
	CoverURL      string   `json:"cover_url"`
	FavoriteCount uint64   `json:"favorite_count"`
	CommentCount  uint64   `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
}
