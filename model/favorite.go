/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 08:58:08
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 08:58:13
 * @FilePath: \GoPath\995_douyin\model\favorite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/model/data"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserId  uint `gorm:"user_id"`
	VideoId uint `gorm:"video_id"`
	// CommentText string   `gorm:"type: text;not null;comment:评论内容"`
	// UserInfo    UserInfo `gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	// Data    interface{} `json:"data"`
}

type FavoriteAction struct {
	UserId     uint   `form:"user_id"`
	Token      string `form:"token" `
	VideoId    uint   `form:"video_id"`
	ActionType int8   `form:"action_type"` // 1 : 点赞， 2 ： 取消点赞
}

type FavoriteList struct {
	UserId uint   `form:"user_id" binding:"required"` // 用户 ID
	Token  string `form:"token" binding:"required"`   // 用户鉴权 token
}

// FeedResponse
type FeedResponse struct {
	Response
	VideoList []data.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

// 用户喜欢视视频回复
type FavoriteListResponse struct {
	Response
	VideoList []data.Video `json:"video_list"`
}

// 获取喜欢列表成功
func GetLikeListSuccess(message string, c *gin.Context, retList []data.Video) {
	c.JSON(http.StatusOK, FavoriteListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  message,
		},
		VideoList: retList,
	})
}

// func GetLikeListSuccess(message string, c *gin.Context) {
// 	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: message})
// }
// 点赞或者取消点赞成功的返回信息
func LikeOperationSuccess(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: message})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: message})
}
