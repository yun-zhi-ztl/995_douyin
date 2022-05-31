package controller

import (
	"995_douyin/config"
	"995_douyin/middleware"
	"995_douyin/model"
	"995_douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list,omitempty"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	userId := Qualify(c).ID
	if user, exist := service.UserInfoQuery(int(userId)); exist {
		//已经登录的情况
		//actionType=1时，是关注操作，=2时，是取关
		actionType, err := strconv.Atoi(c.Query("action_type"))
		if err != nil {
			c.JSON(http.StatusFailedDependency, Response{StatusCode: 0, StatusMsg: "获得对方Id错误!"})
		}
		targetId, _ := strconv.Atoi(c.Query("to_user_id"))
		if actionType == 1 {
			service.Follow(int(user.ID), targetId)
		} else {
			service.Unfollow(int(user.ID), targetId)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "RelationAction success!"})
	} else {
		//还未登录，需要跳转到登录页面
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't (exist)login"})
	}
}

// FollowList
//需要返回UserListResponse格式
func FollowList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	if user, exist := service.UserInfoQuery(userId); exist {
		//查看别人的关注列表
		userList := QueryFolloweeUserList(userId, *user)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	} else {
		//查看自己的关注列表
		var user = Qualify(c)
		userList := QueryFolloweeUserList(userId, user)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	if user, exist := service.UserInfoQuery(userId); exist {
		//查看别人的关注列表
		userList := QueryFollowerUserList(userId, *user)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	} else {
		//查看自己的关注列表
		var user = Qualify(c)
		userList := QueryFollowerUserList(userId, user)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: userList,
		})
	}
}

// QueryFolloweeUserList 查user对应的Followee的UserList
func QueryFolloweeUserList(userId int, user model.User) []User {
	userinfo := service.FindFollowees(int(user.ID))
	//l := len(userinfo)
	//var userList []User
	userList := make([]User, 0, len(userinfo))
	for _, user := range userinfo {
		followCount := service.FolloweeCount(user.Id)
		followerCount := service.FollowerCount(user.Id)
		follow := service.HasFollow(userId, user.Id)
		userList = append(userList, User{
			Id:            int64(user.Id),
			Name:          user.Name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      follow,
		})
	}
	return userList
}

// QueryFollowerUserList 查user对应的Follower的UserList
func QueryFollowerUserList(userId int, user model.User) []User {
	userinfo := service.FindFollowers(int(user.ID))
	//l := len(userinfo)
	//var userList []User
	userList := make([]User, 0, len(userinfo))
	for _, user := range userinfo {
		followCount := service.FolloweeCount(user.Id)
		followerCount := service.FollowerCount(user.Id)
		follow := service.HasFollow(userId, user.Id)
		userList = append(userList, User{
			Id:            int64(user.Id),
			Name:          user.Name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      follow,
		})
	}
	return userList
}

func Qualify(c *gin.Context) model.User {
	jwt, err := middleware.ParserToken(c.Query("token"))
	if err != nil {
		fmt.Errorf("token prase fail!")
	}
	userName := jwt.Username
	var user model.User
	config.DB.Where("Name = ?", userName).Find(&user)
	return user
}
