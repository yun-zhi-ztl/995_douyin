package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/service"
	"github.com/yun-zhi-ztl/995_douyin/utils"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	user_id, _ := utils.ParserToken(token)
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	// 发布评论
	if action_type == "1" {
		comment_text := c.Query("comment_text")
		comment_info, err := service.CreateComment(user_id, video_id, comment_text)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		//log.Println(user_id)
		userinfo, _ := service.QueryUser(user_id, user_id)
		// log.Println(userinfo)
		user := User{
			Id:            uint(user_id),
			Name:          userinfo.Name,
			FollowCount:   int64(userinfo.FollowCount),
			FollowerCount: int64(userinfo.FollowerCount),
			// 此时需要判断IsFollow
			IsFollow: userinfo.IsFollow,
		}
		// 此时要根据comment_info的CreateDate获取CreateDate
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 0, StatusMsg: "success"},
			Comment: Comment{
				Id:         comment_info.ID,
				User:       user,
				Content:    comment_info.CommentText,
				CreateDate: comment_info.CreatedAt.Format("01-02"),
			},
		})
		return
	}
	// 删除评论
	if action_type == "2" {
		comment_id := c.Query("comment_id")
		err := service.DeleteComment(comment_id)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	video_id := c.Query("video_id")
	token := c.Query("token")
	user_id, err := utils.ParserToken(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	comments, err := service.QueryCommentList(video_id)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if len(comments) == 0 {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "success"},
		})
		return
	}
	comment_list := make([]Comment, 0, len(comments))
	for _, comment := range comments {
		userinfo, exist := service.QueryUser(user_id, int(comment.UserId))
		if !exist {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "something error"})
			return
		}
		user := User{
			Id:            uint(userinfo.Id),
			Name:          userinfo.Name,
			FollowCount:   int64(userinfo.FollowCount),
			FollowerCount: int64(userinfo.FollowerCount),
			IsFollow:      userinfo.IsFollow,
		}
		comment_list = append(comment_list, Comment{
			Id:         comment.ID,
			User:       user,
			Content:    comment.CommentText,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "success"},
		CommentList: comment_list,
	})
}
