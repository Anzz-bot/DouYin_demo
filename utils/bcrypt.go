/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 03:10:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 03:10:11
 */
package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// encryption
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Print(err)
	}
	return string(hash)
}

// check pwd
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
