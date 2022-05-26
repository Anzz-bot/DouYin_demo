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
