package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/service"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	// 发布评论
	if action_type == "1" {
		comment_text := c.Query("comment_text")
		_, err := service.CreateComment(user_id, token, video_id, comment_text)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		var userinfo model.UserInfo
		config.DB.Where("id=?", user_id).Find(&userinfo)
		user := User{
			Id:            userinfo.ID,
			Name:          userinfo.UserName,
			FollowCount:   int64(userinfo.FollowCount),
			FollowerCount: int64(userinfo.FollowerCount),
			IsFollow:      userinfo.IsFollow,
		}

		c.JSON(http.StatusOK, CommentResponse{
			BaseResponse: Response{StatusCode: 0},
			ID:           uint(2),
			User:         user,
			CreateDate:   "test-test",
		})
		return
	}
	// 删除评论
	if action_type == "2" {
		comment_id := c.Query("comment_id")
		err := service.DeleteComment(comment_id)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	}

	// if _, exist := UsersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
