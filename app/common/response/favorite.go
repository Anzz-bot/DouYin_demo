package response

type FavoriteListResponse struct {
	Response
	VideoList []*VideoAuthorApi `json:"video_list"`
}
