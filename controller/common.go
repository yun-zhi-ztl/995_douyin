/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 10:42:37
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:07:13
 * @FilePath: \GoPath\995_douyin\controller\common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         uint   `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// CommentResponse 评论操作响应内容
type CommentResponse struct {
	Response
	Comment Comment
}

// CommentListResponse 评论列表响应内容
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

func Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: msg})
}

func Failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: msg})
}
