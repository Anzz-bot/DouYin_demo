package request

type Relation struct {
	UserId     uint64 `form:"user_id" json:"user_id" `
	ToUserId   uint64 `form:"to_user_id" json:"to_user_id" `
	Token      string `form:"token" json:"token" `
	ActionType int32  `form:"action_type" json:"action_type" `
}
