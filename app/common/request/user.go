/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 01:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 01:14:50
 */
package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Login struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误信息
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名称不能为空",
		"Password.required": "用户密码不能为空",
	}
}

// 自定义错误信息
func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名称不能为空",
		"Password.required": "用户密码不能为空",
	}
}
