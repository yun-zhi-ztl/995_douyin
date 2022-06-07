/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 10:42:37
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:14:29
 * @FilePath: \GoPath\995_douyin\controller\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/service"
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
	User *service.UserInfo `json:"user"`
}

// 注册response：UserLoginResponse
type UserRegisterResponse struct {
	Response
	UserID int    `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// Register
//  @Description: 用户注册接口
//  @param c *gin.Context
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	registerinfo := service.Register(username, password)
	if registerinfo.Err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: registerinfo.Err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 0},
			UserID:   registerinfo.Userid,
			Token:    registerinfo.Token,
		})
	}
}

// Login
//  @Description: 用户登录接口
//  @param c *gin.Context
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// username, _ := c.GetPostForm("username")
	// password, _ := c.GetPostForm("password")
	loginfo := service.Login(username, password)
	if loginfo.Err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: loginfo.Err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{StatusCode: 0},
			UserID:   loginfo.Userid,
			Token:    loginfo.Token,
		})
	}
}

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	token := c.Query("token")
	user, err := service.QueryUserInfo(token, userid)
	if err != nil {
		//if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}
