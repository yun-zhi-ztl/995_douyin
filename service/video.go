/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 09:00:51
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:15:03
 * @FilePath: \GoPath\995_douyin\service\video.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/model/data"
)

func GetLikeVideoList(videolist []model.Video) ([]data.Video, error) {
	// 需要更新 is_follower 和 is_favorite 这两个参数
	// 因为需要判断当前用户是不是这些视频作者的关注者
	// var responsevideolist []response.Video
	responsevideolist := make([]data.Video, len(videolist))
	for i, video := range videolist {
		responsevideolist[i] = data.Video{
			Id: video.ID,
			Author: data.User{
				Id:            video.UserId,
				Name:          video.Author.UserName,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      true,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
	}
	return responsevideolist, nil
}
