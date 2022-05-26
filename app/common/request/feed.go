/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 02:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 02:53:50
 */
package request

type Feed struct {
	LatestTime int64  `form:"latest_time" json:"latest_time" `
	Token      string `form:"token" json:"token" `
}
