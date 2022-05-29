package service

import (
	"errors"
	"strconv"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
)

// 增加评论
func CreateComment(userid, token, videoid, comment_text string) (*model.Comment, error) {
	user_id, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		return nil, err
	}
	video_id, err := strconv.ParseUint(videoid, 10, 64)
	if err != nil {
		return nil, err
	}
	// token_id, err := middleware.ParserToken(token)
	// if err != nil {
	// 	return nil, err
	// }
	// if token_id != int(user_id) {
	// 	return nil, errors.New("user is didn't login")
	// }
	// comment := model.CreateNewComment(uint(user_id), uint(video_id), comment_text)
	comment := &model.Comment{
		UserId:      uint(user_id),
		VideoID:     uint(video_id),
		CommentText: comment_text,
	}
	err = config.DB.Create(&comment).Error
	if err != nil {
		// return nil, errors.New("error in create user")
		return nil, err
	}
	// 更新video模型中的comment总数
	return comment, nil
}

// 删除评论
func DeleteComment(commentid string) error {
	comment_id, err := strconv.ParseUint(commentid, 10, 64)
	if err != nil {
		return err
	}
	comment := model.QuerComment(uint(comment_id))
	if comment == nil {
		return errors.New("comment is not exist")
	}
	// fmt.Println(comment)
	// 软删除
	err = model.DeleteComment(comment)
	return err
}
