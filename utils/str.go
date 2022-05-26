/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 01:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-23 21:40:51
 */
package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func StrToUint64(str string) uint64 {
	res, _ := strconv.ParseUint(str, 10, 64)
	return res
}

func StrToInt64(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func StrToUint32(str string) uint64 {
	res, _ := strconv.ParseUint(str, 10, 32)
	return res
}

func StrToInt32(str string) int32 {
	res, _ := strconv.ParseInt(str, 10, 32)
	return int32(res)
}
