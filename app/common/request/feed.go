package request

type Feed struct {
	LatestTime int64  `form:"latest_time" json:"latest_time" `
	Token      string `form:"password" json:"password" `
}
