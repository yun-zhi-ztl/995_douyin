/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-29 07:07:50
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:32:41
 * @FilePath: \GoPath\995_douyin\service\comment.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"errors"
	"strconv"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
)

// 增加评论
/**
 * @description:
 * @param {*} userid
 * @param {*} token
 * @param {*} videoid
 * @param {string} comment_text
 * @return {*}
 */
func CreateComment(userid int, videoid, comment_text string) (*model.Comment, error) {
	// user_id, err := strconv.ParseUint(userid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
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
	comment := model.CreateNewComment(uint(userid), uint(video_id), comment_text)
	err = config.DB.Create(&comment).Error
	if err != nil {
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

// 获取评论列表
func QueryCommentList(videoid string) ([]model.Comment, error) {
	var comment_list []model.Comment
	video_id, err := strconv.ParseUint(videoid, 10, 64)
	if err != nil {
		return nil, err
	}
	config.DB.Where("video_id=?", video_id).Find(&comment_list)
	if len(comment_list) == 0 {
		return nil, errors.New("this video has zero comment")
	}
	return comment_list, nil
}
