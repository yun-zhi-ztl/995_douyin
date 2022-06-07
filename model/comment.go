/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 09:14:54
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:52:40
 * @FilePath: \GoPath\995_douyin\model\comment.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

func CreateNewComment(user_id, video_id uint, comment_text string) (*Comment, error) {
	comment := &Comment{
		UserId:      user_id,
		VideoID:     video_id,
		CommentText: comment_text,
	}
	tx := config.DB.Begin()
	err := config.DB.Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	err = config.DB.Model(&Video{}).Where("id=?", video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return comment, nil
}

func QuerComment(comment_id uint) *Comment {
	var comment Comment
	config.DB.Where("id=?", comment_id).Find(&comment)
	return &comment
}

func DeleteComment(del_comment *Comment) error {
	tx := config.DB.Begin()
	err := config.DB.Delete(&del_comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	err = config.DB.Model(&Video{}).Where("id = ?", del_comment.VideoID).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
