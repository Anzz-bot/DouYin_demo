/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/Anzz-bot/DouYin_demo/utils"
	"strconv"
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

	user = models.User{Name: params.Name, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error

	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("name = ?", params.Name).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	global.NowUserID = user.ID.ID
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user *models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
