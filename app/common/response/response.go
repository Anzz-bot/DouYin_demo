/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 03:10:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 03:10:11
 */
package response

import (
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 响应结构体
type Response struct {
	ErrorCode int         `json:"error_code"` // 自定义错误码
	Data      interface{} `json:"data"`       // 数据
	Message   string      `json:"message"`    // 信息
}

type CodeResponse struct {
	StatusCode int    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

func NewResponse(code int, msg string) CodeResponse {
	return CodeResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}
