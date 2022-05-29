package service

import (
	"errors"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/middleware"
	"github.com/yun-zhi-ztl/995_douyin/model"
)

type RegisterInfo struct {
	Err    error
	Token  string
	Userid int
}

// 用户注册
func Register(username, password string) *RegisterInfo {
	// 将user写进数据库
	user, err := model.CreateNewUserSingleton(username, password)
	if err != nil {
		return &RegisterInfo{
			Err: err,
		}
	}
	// 此处需要考虑数据库线程安全
	if config.DB.Create(&user).Error != nil {
		return &RegisterInfo{
			Err: errors.New("error in create user"),
		}
	}
	// 生成token
	token, err := middleware.CreateJwtToken(user.ID)
	if err != nil {
		return &RegisterInfo{
			Err: errors.New("error in token generation"),
		}
	}
	return &RegisterInfo{
		Err:    nil,
		Userid: int(user.ID),
		Token:  token,
	}
}

type LoginInfo struct {
	Err    error
	Token  string
	Userid int
}

// 用户登录
func Login(username, password string) *LoginInfo {
	// 用户名和密码不能为空
	if len(username) == 0 || len(password) == 0 {
		return &LoginInfo{
			Err: errors.New("user_name and password cannot be empty"),
		}
	}
	// 检查是否已经注册
	var user model.UserInfo
	config.DB.Where("user_name = ?", username).Find(&user)
	if user.ID == 0 {
		return &LoginInfo{
			Err: errors.New("the user not registered, please registe in"),
		}
	}
	// 判断username
	if user.Password != password {
		return &LoginInfo{
			Err: errors.New("the password is error"),
		}
	}
	// 生成token
	token, err := middleware.CreateJwtToken(user.ID)
	if err != nil {
		return &LoginInfo{
			Err: errors.New("error in token generation"),
		}
	}
	return &LoginInfo{
		Err:    nil,
		Userid: int(user.ID),
		Token:  token,
	}
}

type UserInfo struct {
	Id            int
	Name          string
	FollowCount   int
	FollowerCount int
	IsFollow      bool
}

func UserQue(token string) (*UserInfo, bool) {
	userid, err := middleware.ParserToken(token)
	userinfo := &UserInfo{
		Id:            userid,
		Name:          "test",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	}
	if err == nil {
		return userinfo, true
	} else {
		return nil, false
	}

}
