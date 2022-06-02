/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-25 00:45:20
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 10:24:52
 * @FilePath: \GoPath\995_douyin\model\user.go
 * @Description: 用户数据库表设计以及数据库相关操作
 */
package model

import (
	"errors"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"gorm.io/gorm"
)

/**
 * @description: 用户数据库表设计
 * @return {*}
 */
type UserInfo struct {
	gorm.Model
	UserName      string `gorm:"varchar(32);not null;unique;comment:用户名称"`
	Password      string `gorm:"varchar(32);not null;comment:用户密码"`
	FollowCount   int    `gorm:"default:0;not null;comment:关注总数"`
	FollowerCount int    `gorm:"default:0;not null;comment:粉丝总数"`
	IsFollow      bool   `gorm:"-"`
}

/**
 * @description: 创建一个用户表
 * @param {string} username：用户姓名
 * @param {string} password：用户密码
 * @return *UserInfo, error：新建的用户信息/错误信息
 */
func CreateNewUserSingleton(username, password string) (*UserInfo, error) {
	// 用户名和密码不能为空
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("username and password cannot be empty")
	}
	// 检查是否已经注册
	var user UserInfo
	config.DB.Model(&UserInfo{}).Where("user_name=?", username).Find(&user)
	if user.ID != 0 {
		return nil, errors.New("the user has already registered, please log in")
	}
	return &UserInfo{
		UserName: username,
		Password: password,
	}, nil
}

func QueryUserInfo(user_id uint) (*UserInfo, error) {
	var user UserInfo
	config.DB.Model(&UserInfo{}).Where("id=?", user_id).Find(&user)
	if user.ID == 0 {
		return nil, errors.New("the user doesn't exist")
	}
	return &user, nil
}
