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

func Register(username, password string) *RegisterInfo {
	// 用户名和密码不能为空
	if len(username) == 0 || len(password) == 0 {
		return &RegisterInfo{
			Err: errors.New("username and password cannot be empty"),
		}
	}
	// 检查是否已经注册
	var user model.User
	config.DB.Where("Name = ?", username).Find(&user)
	if user.ID != 0 {
		return &RegisterInfo{
			Err: errors.New("the user has already registered, please log in"),
		}
	}
	// 将user写进数据库
	newUser := model.User{
		Name:     username,
		Password: password,
	}
	config.DB.Create(&newUser)
	// 缺少对token的处理
	// token, err := middleware.CreateJwtToken(username, password)
	token, err := middleware.CreateJwtToken1(int(newUser.ID))
	if err != nil {
		return &RegisterInfo{
			Err: errors.New("error in token generation"),
		}
	}
	return &RegisterInfo{
		Err:    nil,
		Userid: int(newUser.ID),
		Token:  token,
	}
}

type LoginInfo struct {
	Err    error
	Token  string
	Userid int
}

func Login(username, password string) *LoginInfo {
	token, err := middleware.CreateJwtToken(username, password)
	if err != nil {
		return &LoginInfo{
			Err: errors.New("error in token generation"),
		}
	}
	// 用户名和密码不能为空
	if len(username) == 0 || len(password) == 0 {
		return &LoginInfo{
			Err: errors.New("username and password cannot be empty"),
		}
	}
	// 检查是否已经注册
	var user model.User
	config.DB.Where("Name = ?", username).Find(&user)
	if user.ID == 0 {
		return &LoginInfo{
			Err: errors.New("the user not registered, please registe in"),
		}
	}
	if user.Password != password {
		return &LoginInfo{
			Err: errors.New("the password is error"),
		}
	}
	// 缺少对token的处理

	return &LoginInfo{
		Err:    nil,
		Userid: int(user.ID),
		Token:  token,
	}
}
