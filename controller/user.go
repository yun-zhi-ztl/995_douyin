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
	// 可通过数据库查询
	// if user, exist := usersLoginInfo[token]; exist {
	// } else {
	// c.JSON(http.StatusOK, UserLoginResponse{
	// 	Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// })
	// }
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

// // token回应：测试处理
// type UserToken struct {
// 	Response
// 	UserName string `json:"username,omitempty"`
// 	Password string `json:"password,omitempty"`
// }

// // 解析token测试
// func Token(c *gin.Context) {
// 	token, _ := c.GetPostForm("token")
// 	claims, err := middleware.ParserToken(token)
// 	if err != nil {
// 		c.JSON(http.StatusOK, UserRegisterResponse{
// 			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, UserToken{
// 			Response: Response{StatusCode: 0},
// 			UserName: claims.Username,
// 			Password: claims.Userpassword,
// 		})
// 	}
// }
