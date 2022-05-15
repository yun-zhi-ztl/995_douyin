package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]User{}

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
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

}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
