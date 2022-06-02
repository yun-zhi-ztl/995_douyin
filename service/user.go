package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/utils"
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
	token, err := utils.CreateJwtToken(user.ID)
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
	token, err := utils.CreateJwtToken(user.ID)
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

func QueryUserInfo(token, userid string) (*UserInfo, error) {
	user_id, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		return nil, err
	}
	token_id, err := utils.ParserToken(token)
	if err != nil {
		return nil, err
	}
	if user_id == 0 {
		user, err := model.QueryUserInfo(uint(token_id))
		if err != nil {
			return nil, err
		}
		userinfo := &UserInfo{
			Id:            int(user.ID),
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		return userinfo, err
	}
	user, err := model.QueryUserInfo(uint(user_id))
	if err != nil {
		return nil, err
	}
	userinfo := &UserInfo{
		Id:            int(user.ID),
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      HasFollow(token_id, int(user_id)),
	}
	return userinfo, err
}

func QueryUser(user_id, to_user_id int) (*UserInfo, bool) {
	user, err := model.QueryUserInfo(uint(to_user_id))
	if err != nil {
		return nil, false
	}
	fmt.Println(user)
	userinfo := &UserInfo{
		Id:            int(user.ID),
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      HasFollow(user_id, to_user_id),
	}
	return userinfo, true
}

func UserInfoQuery(userId int) (*model.UserInfo, bool) {
	var user model.UserInfo
	config.DB.Where("Id = ?", userId).Find(&user)
	if user.ID == 0 {
		return &user, false
	} else {
		return &user, true
	}
}
