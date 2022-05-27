package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTimestamp := c.Query("latest_time")
	startTime := ""
	if lastTimestamp != "" {
		if parseIntRes, parseIntErr := strconv.ParseInt(lastTimestamp, 10, 64); parseIntErr == nil {
			startTime = time.Unix(parseIntRes/1000, 0).Format("2006-01-02 15:04:05")
		}
	}
	feedVideoList, getFeedErr := videoService.Feed(startTime)
	if getFeedErr != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  getFeedErr.Error(),
		})
		return
	}
	lenFeedVideoList := len(feedVideoList)
	if lenFeedVideoList == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	userId := c.GetInt("userID")                                   // 没有设置会取类型默认值
	nextTime := feedVideoList[lenFeedVideoList-1].CreatedAt.Unix() // 防止后续的排序影响
	videoList := make([]Video, 0, lenFeedVideoList)
	for _, video := range feedVideoList {
		videoList = append(videoList, Video{
			Id: int64(video.ID),
			Author: User{
				Id:            int64(video.User.ID),
				Name:          video.User.UserName,
				FollowCount:   int64(video.User.FollowerCount),
				FollowerCount: int64(video.User.FollowerCount),
				IsFollow:      IsFollow(uint(userId), video.User.ID),
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    IsFavorite(video.User.ID, video.ID),
		})
	}
	if userId != 0 {
		// 如果用户登录 则优先展示当前登录用户关注的作者的视频
		// TODO: 兼容1.7及以下
		//sort.Slice(videoList, func(i, j int) bool { return videoList[i].Author.IsFollow || videoList[j].Author.IsFollow }) // g0 1.8+支持
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 1, StatusMsg: "success"},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}

// IsFavorite
//  @Description: 判断用户是否点赞当前视频
//  @param userId uint 当前登录用户ID
//  @param videoId uint	当前视频ID
//  @return bool 未登录则直接返回false
func IsFavorite(userId, videoId uint) bool {
	if userId == 0 {
		return false
	}
	return true
	//return favoriteService.IsFavorite(userId, videoId) // 点赞操作业务应该提供该接口
}

// IsFollow
//  @Description: 判断登录用户是否关注了视频作者
//  @param fromUserId uint 登录用户
//  @param toUserId uint 视频作者
//  @return bool 未登录直接返回false
func IsFollow(fromUserId, toUserId uint) bool {
	if fromUserId == 0 {
		return false
	}
	return true
	//return relationService.IsFollow(fromUserId, toUserId) // 关系操作业务应该提供该接口
}
