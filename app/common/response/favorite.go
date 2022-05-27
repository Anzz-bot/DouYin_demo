/*
 * @Author: alexander.huang
 * @Date:   2022-05-26 12:44:17
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-26 12:44:17
 */
package response

type FavoriteListResponse struct {
	Response
	VideoList []*VideoAuthorApi `json:"video_list"`
}
