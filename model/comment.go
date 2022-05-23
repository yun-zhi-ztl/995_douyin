package model

import "gorm.io/gorm"

// 评论id和创建时间由gorm.Model自动生成
type Comment struct {
	gorm.Model
	UserId      uint   `gorm:"default:0,not null;comment:创作用户ID"`
	VideoID     uint   `gorm:"default:0,not null;comment:视频ID"`
	CommentText string `gorm:"varchar(32);not null;comment:评论内容"`
}
