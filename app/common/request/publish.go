/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 22:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 22:53:50
 */
package request

import (
	"mime/multipart"
)

type VideoUpload struct {
	Data  *multipart.FileHeader `form:"data" json:"data" binding:"required"`
	Token string                `form:"token" json:"token"`
	Title string                `form:"title" json:"title" binding:"required"`
}
type PublishList struct {
	UserId string `form:"user_id" json:"user_id" `
	Token  string `form:"token" json:"token" `
}

func (videoUpload VideoUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"title.required": "标题不能为空",
		"data.required":  "请选择视频",
	}
}
