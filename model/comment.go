package model

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
	"gorm.io/gorm"
)

// 评论id和创建时间由gorm.Model自动生成
type Comment struct {
	gorm.Model
	UserId      uint     `gorm:"not null;comment:创作用户ID"`
	VideoID     uint     `gorm:"not null;comment:视频ID"`
	CommentText string   `gorm:"type: text;not null;comment:评论内容"`
	UserInfo    UserInfo `gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
}

func CreateNewComment(user_id, video_id uint, comment_text string) *Comment {
	return &Comment{
		UserId:      user_id,
		VideoID:     video_id,
		CommentText: comment_text,
	}
}

func QuerComment(comment_id uint) *Comment {
	var comment Comment
	config.DB.Where("id=?", comment_id).Find(&comment)
	return &comment
}

func DeleteComment(del_comment *Comment) error {
	return config.DB.Delete(&del_comment).Error
}
