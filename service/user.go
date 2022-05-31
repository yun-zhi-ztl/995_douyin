package service

import (
	"errors"
	"strconv"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/middleware"
	"github.com/yun-zhi-ztl/995_douyin/model"
)

type RegisterInfo struct {
	Err    error
	Token  string
	Userid int
}

/**
 * @description: 用户注册
 * @param {string} username：用户姓名
 * @param {string} password：用户密码
 * @return {*} RegisterInfo：注册信息{错误信息、token、用户id}
 */
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

/**
 * @description：用户登录
 * @param {*} username：用户名
 * @param {string} password：密码
 * @return {*} 用户登录信息
 */
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

func QueryUserInfo(token, userid string) (*UserInfo, bool) {
	user_id, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		return nil, false
	}
	token_id, err := middleware.ParserToken(token)
	if err != nil {
		return nil, false
	}
	if user_id == 0 {
		user, err := model.QueryUserInfo(uint(token_id))
		if err != nil {
			return nil, false
		}
		userinfo := &UserInfo{
			Id:            int(user.ID),
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		return userinfo, true
	}
	user, err := model.QueryUserInfo(uint(user_id))
	if err != nil {
		return nil, false
	}
	userinfo := &UserInfo{
		Id:            int(user.ID),
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      HasFollow(token_id, int(user_id)),
	}
	return userinfo, true
}

// !需要改动
func HasFollow(token_id, user_id int) bool {
	return false
}

func QueryUser(user_id, to_user_id int) (*UserInfo, bool) {
	user, err := model.QueryUserInfo(uint(to_user_id))
	if err != nil {
		return nil, false
	}
	userinfo := &UserInfo{
		Id:            int(user.ID),
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      HasFollow(user_id, to_user_id),
	}
	return userinfo, true
}
