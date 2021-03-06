/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-15 22:11:28
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:07:16
 * @FilePath: \GoPath\995_douyin\controller\feed.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/utils"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/service"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

var videoService service.VideoService

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTimestamp := c.Query("latest_time")
	startTime := ""
	if lastTimestamp != "" {
		if parseIntRes, parseIntErr := strconv.ParseInt(lastTimestamp, 10, 64); parseIntErr == nil {
			startTime = time.Unix(parseIntRes, 0).Format("2006-01-02 15:04:05")
		}
	} else {
		startTime = time.Now().Format("2006-01-02 15:04:05")
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
	token := c.Query("token")
	var userId int
	var err error
	if token != "" {
		userId, err = utils.ParserToken(token)
		if err != nil {
			c.JSON(http.StatusOK, UserRegisterResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
		}
	} else {
		userId = 0
	}
	// userId := c.GetInt("userID")                                // 没有设置会取类型默认值
	nextTime := feedVideoList[lenFeedVideoList-1].CreatedAt.Unix() // 防止后续的排序影响
	videoList := make([]Video, 0, lenFeedVideoList)
	for _, video := range feedVideoList {
		videoList = append(videoList, Video{
			Id: int64(video.ID),
			Author: User{
				Id:            video.Author.ID,
				Name:          video.Author.UserName,
				FollowCount:   int64(video.Author.FollowerCount),
				FollowerCount: int64(video.Author.FollowerCount),
				IsFollow:      IsFollow(uint(userId), video.Author.ID),
			},
			PlayUrl:       config.ServerDomain + video.PlayUrl,
			CoverUrl:      config.ServerDomain + video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    IsFavorite(video.ID, uint(userId)),
			Title:         video.Title,
		})
	}
	if userId != 0 {
		sort.Slice(videoList, func(i, j int) bool { return videoList[i].Author.IsFollow || videoList[j].Author.IsFollow }) // g0 1.8+支持
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "success"},
		NextTime:  nextTime,
		VideoList: videoList,
	})
}

// IsFavorite
//  @Description: 判断用户是否点赞当前视频
//  @param userId uint 当前登录用户ID
//  @param videoId uint	当前视频ID
//  @return bool 未登录则直接返回false
func IsFavorite(videoId, userId uint) bool {
	return favoriteService.IsFavorite(userId, videoId)
	//return favoriteService.IsFavorite(userId, videoId) // 点赞操作业务应该提供该接口
}

// IsFollow
//  @Description: 判断登录用户是否关注了视频作者
//  @param fromUserId uint 登录用户
//  @param toUserId uint 视频作者
//  @return bool 未登录直接返回false
func IsFollow(fromUserId, toUserId uint) bool {
	return service.HasFollow(int(fromUserId), int(toUserId)) // 关系操作业务应该提供该接口
}
