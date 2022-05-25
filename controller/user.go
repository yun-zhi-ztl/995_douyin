package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]User{}

type UserLoginResponse struct {
	status_code int
	status_msg  string
	user_id     int
	token       string
}

// Register
//  @Description: 用户注册接口
//  @param c *gin.Context
func Register(c *gin.Context) {

}

// Login
//  @Description: 用户登录接口
//  @param c *gin.Context
func Login(c *gin.Context) {
	response := gin.H{
		"status_code": 0,
		"status_msg":  "登录成功",
		"user_id":     1111111,
		"token":       "sdjsdkdiosfksafafhsus",
	}
	user := User{
		Name: "SSSSSSS",
	}
	UsersLoginInfo["sdjsdkdiosfksafafhsus"] = user
	c.JSON(http.StatusOK, response)
}

func UserInfo(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	user := User{
		Id:            int64(user_id),
		Name:          "sssssss",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	response := gin.H{
		"status_code": 0,
		"status_msg":  "登录成功",
		"user":        user,
	}
	c.JSON(http.StatusOK, response)
}
