/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 03:10:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 03:10:11
 */
package response

import "github.com/Anzz-bot/DouYin_demo/app/models"

const (
	CodeUserNotFound = 1000 + iota
	CodeUserLoginFailed
	CodeUserAlreadyExists
	CodeUserRegisterFailed
	CodeUserNameInvalid
	CodeUserPasswordInvalid
	CodeUserIDInvalid
	CodeUserTokenInvalid
	CodeSuccess
)

var CodeUserMessages = map[int]string{
	CodeSuccess:             "",
	CodeUserNotFound:        "用户不存在",
	CodeUserLoginFailed:     "用户名或密码错误",
	CodeUserAlreadyExists:   "用户名已存在",
	CodeUserRegisterFailed:  "注册失败",
	CodeUserNameInvalid:     "不合法的用户名",
	CodeUserPasswordInvalid: "不合法的密码",
	CodeUserIDInvalid:       "不合法的用户ID",
	CodeUserTokenInvalid:    "Token无效",
}

// UserRegisterResponse 用户注册响应
type UserRegisterResponse struct {
	CodeResponse
	UserID models.ID `json:"user_id"` // 用户ID
	Token  string    `json:"token"`   // 用户鉴权token
}

func BuildUserRegisterResponse(code int, user *models.User, token string) *UserRegisterResponse {
	res := &UserRegisterResponse{}
	if code != CodeSuccess {
		res.CodeResponse = NewResponse(code, CodeUserMessages[code])
		return res
	} else {
		res.CodeResponse = NewResponse(CodeSuccess, CodeUserMessages[CodeSuccess])
	}
	res.UserID = user.ID
	res.Token = token
	return res
}
