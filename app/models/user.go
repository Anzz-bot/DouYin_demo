/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 22:20:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 22:20:50
 */
package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Token    string `json:"token"    gorm:"not null;comment:鉴权token"`
	Timestamps
	SoftDeletes
}
