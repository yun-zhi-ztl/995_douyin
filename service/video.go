// Package service
// @author ufec https://github.com/ufec
// @date 2022/5/27
package service

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
)

type VideoService struct {
}

// Feed
//  @Description: 获取视频流
//  @receiver s *VideoService
//  @param startTime string 起始时间
//  @return []model.Video
//  @return error
func (s *VideoService) Feed(startTime string) ([]model.Video, error) {
	var videoList []model.Video
	err := config.DB.Where("created_at <= ?", startTime).Preload("User").Order("created_at DESC").Limit(30).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
