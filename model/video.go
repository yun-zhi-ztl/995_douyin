package model

import "gorm.io/gorm"

// 视频id和创建时间由gorm.Model自动生成
type Video struct {
	gorm.Model
	// 通过其获得视频列表中的用户信息
	UserId        uint   `gorm:"default:0,not null;comment:创作用户ID"`
	PlayUrl       string `gorm:"varchar(32);not null;comment:视频播放地址"`
	CoverUrl      string `gorm:"varchar(32);not null;comment:视频封面地址"`
	Title         string `gorm:"varchar(32);not null;comment:视频标题"`
	FavoirteCount int    `gorm:"default:0,not null;comment:视频点赞总数"`
	CommentCount  int    `gorm:"default:0,not null;comment:视频评论总数"`
	IsFavorite    bool   `gorm:"default:false,not null;comment:是否点赞"`
}
