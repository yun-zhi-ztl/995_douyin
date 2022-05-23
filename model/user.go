package model

import (
	"errors"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"gorm.io/gorm"
)

// 用户id即gorm.Model自动生成
type UserInfo struct {
	gorm.Model
	UserName      string `gorm:"varchar(32);not null;unique;comment:用户名称"`
	Password      string `gorm:"varchar(32);not null;comment:用户密码"`
	FollowCount   int    `gorm:"default:0,not null;comment:关注总数"`
	FollowerCount int    `gorm:"default:0,not null;comment:粉丝总数"`
	IsFollow      bool   `gorm:"default:false,not null;comment:是否关注"`
}

func CreateNewUserSingleton(username, password string) (*UserInfo, error) {
	// 用户名和密码不能为空
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("username and password cannot be empty")
	}
	// 检查是否已经注册
	var user UserInfo
	config.DB.Where("Name = ?", username).Find(&user)
	if user.ID != 0 {
		return nil, errors.New("the user has already registered, please log in")
	}
	return &UserInfo{
		UserName:      username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}, nil
}
