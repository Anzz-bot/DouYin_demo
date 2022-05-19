/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 20:14:50
 */
package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("name = ?", params.Name).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("用户已存在")
		return
	}
	user = models.User{Name: params.Name, Password: utils.BcryptMake([]byte(params.Password)), Token: params.Name}
	err = global.App.DB.Create(&user).Error
	return
}
