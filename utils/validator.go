/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 03:10:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 03:10:11
 */
package utils

//Define validation rules. Other validation rules will be stored here

func CheckName(name string) bool {
	if len(name) > 32 || len(name) < 0 {
		return false
	}
	return true
}

func CheckPasswd(passwd string) bool {
	if len(passwd) > 32 || len(passwd) < 0 {
		return false
	}
	return true
}
