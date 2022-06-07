/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-25 00:45:20
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 07:52:19
 * @FilePath: \GoPath\995_douyin\model\video.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
	"gorm.io/gorm"
)

// 视频id和创建时间由gorm.Model自动生成
type Video struct {
	gorm.Model
	// 通过其获得视频列表中的用户信息
	UserId        uint     `gorm:"default:0;not null; comment:创作用户ID"`
	PlayUrl       string   `gorm:"varchar(32);not null;comment:视频播放地址"`
	CoverUrl      string   `gorm:"varchar(32);not null;comment:视频封面地址"`
	Title         string   `gorm:"varchar(32);not null;comment:视频标题"`
	FavoriteCount int      `gorm:"default:0;not null;comment:视频点赞总数"`
	CommentCount  int      `gorm:"default:0;not null;comment:视频评论总数"`
	IsFavorite    bool     `gorm:"default:false;not null;comment:是否点赞"`
	Author        UserInfo `gorm:"foreignKey:UserId; references:ID; comment:视频作者"`
}

func (v Video) TableName() string {
	return "douyin_videos"
}

func (v Video) Create() error {
	return config.DB.Create(&v).Error
}

func (v Video) QueryByUserID(userId int) (*[]Video, error) {
	var videos []Video
	// err := config.DB.Where("user_id = ?", userId).Preload("Author").Find(&videos).Error
	err := config.DB.Where(&Video{UserId: uint(userId)}).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return &videos, err
}
