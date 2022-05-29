/*
 * @Author: alexander.huang
 * @Date:   2022-05-25 02:53:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-25 02:53:50
 */
package services

import (
	"errors"
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/global"
	"time"
)

type feedService struct {
}

var FeedService = new(feedService)

func (feedService *feedService) Feed(params request.Feed) (err error, response response.FeedResponse) {
	latestTime, feedList, err := feedService.GetFeedList(time.UnixMilli(params.LatestTime))
	if err != nil {
		global.App.Log.Info(err.Error())
		return
	}
	response.VideoList = feedList
	response.NextTime = latestTime.UnixMilli()
	return
}

//todo： ffmpeg 实现抽帧

func (feedService *feedService) GetFeedList(unixTime time.Time) (time.Time, []*response.VideoAuthorApi, error) {
	var videoAuthorApi []*response.VideoAuthorApi
	err := global.App.DB.Raw("SELECT\n    v.ID AS id,\n    u.id AS author_id,\n    u.name AS author_name,\n    u.follow_count AS author_follow_count,\n    u.follower_count AS author_follower_count,\n    false AS author_is_follow,\n    v.play_url AS play_url,\n    v.cover_url AS cover_url,\n    v.favorite_count AS favorite_count,\n    v.comment_count AS comment_count,\n    false AS is_favorite\nFROM videos v\nLEFT JOIN users u ON v.author_id=u.id\nWHERE v.created_at < ?\nLIMIT 30;", unixTime).Scan(&videoAuthorApi).Error
	if err != nil {
		return unixTime, nil, err
	}
	if len(videoAuthorApi) == 0 {
		return unixTime, nil, errors.New("no video found")
	}
	lastVideoID := videoAuthorApi[len(videoAuthorApi)-1].ID
	type CreatedAt struct {
		CreatedAt time.Time `json:"created_at"`
	}
	lastVideoCreatedAt := &CreatedAt{}
	err = global.App.DB.Raw("SELECT created_at FROM videos WHERE id = ?", lastVideoID).Scan(lastVideoCreatedAt).Error
	if err != nil {
		return unixTime, nil, err
	}
	return lastVideoCreatedAt.CreatedAt, videoAuthorApi, nil
}
