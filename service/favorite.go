package service

import (
	"log"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"gorm.io/gorm"
)

type FavoriteService struct{}

func (like *FavoriteService) IsFavorite(userId, videoId uint) bool {
	var count int64
	err := config.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error
	if err != nil {
		log.Println("IsFavorite error in ./service/favorite.go")
	}
	if count != 0 {
		return true
	}
	return false
}

func (like *FavoriteService) LikeOperation(likeinfo model.FavoriteAction) error {
	actiontype := likeinfo.ActionType
	var err error
	if actiontype == 1 {
		err = like.UserLike(likeinfo)
	} else {
		err = like.UserUnLike(likeinfo)
	}
	return err
}

func (like *FavoriteService) UserLike(likeinfo model.FavoriteAction) error {
	actiontype := likeinfo.ActionType
	// num, err := favoriteOperation.GetFavoriteCount(likeinfo.UserId, likeinfo.VideoId)
	var count int64
	err := config.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", likeinfo.UserId, likeinfo.VideoId).Count(&count).Error
	if err != nil {
		log.Println("UserLike error in ./service/favorite.go")
	}

	if count == 0 && actiontype == 1 {
		// 需要点赞操作
		// 开启事务
		tx := config.DB.Begin()
		var like model.Favorite
		like.UserId = likeinfo.UserId
		like.VideoId = likeinfo.VideoId

		err := tx.Create(&like).Error
		if err != nil {
			log.Println("LikeOperation error in service/favorite.go")
			tx.Rollback() // 点赞失败则回滚点赞操作
			return err
		} else {
			tx.Commit() // 点赞成功提交这个订单，之后可以在查找表的时候获取这个点赞数量字段
		}

		// 更新 video 表单对应的字段
		err = config.DB.Model(&model.Video{}).Where("id = ?", likeinfo.VideoId).Update("is_favorite", true).Error
		if err != nil {
			log.Println("UpdateVideoAfterLikeOperation is_favorite error in ./service/favorite.go")
		}
		// 更新 Video 数据结构中的 favorite_count 字段
		err = config.DB.Model(&model.Video{}).Where("id = ?", likeinfo.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			log.Println("UpdateVideoAfterLikeOperation favoritr_count error in ./service/favorite.go")
		}
	}
	return nil
}

func (like *FavoriteService) UserUnLike(likeinfo model.FavoriteAction) error {

	actiontype := likeinfo.ActionType
	var count int64
	err := config.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", likeinfo.UserId, likeinfo.VideoId).Count(&count).Error
	if err != nil {
		log.Println("UserLike error in ./service/favorite.go")
	}
	if count != 0 && actiontype == 2 {
		// 需要进入数据库中，删除这个点赞操作
		tx := config.DB.Begin()
		var unlike model.Favorite
		unlike.UserId = likeinfo.UserId
		unlike.VideoId = likeinfo.VideoId
		err := tx.Unscoped().Where("user_id = ? and video_id = ?", likeinfo.UserId, likeinfo.VideoId).Delete(&unlike).Error
		if err != nil {
			log.Println("UnLikeOperation error in database/favorite.go")
			tx.Rollback() // 取消点赞失败则回滚取消这次的点赞操作
			return err
		} else {
			tx.Commit()
		}
		err = config.DB.Model(&model.Video{}).Where("id = ?", likeinfo.VideoId).Update("is_favorite", false).Error
		if err != nil {
			log.Println("UpdateVideoAfterUnLikeOperation is_favorite error in ./service/favorite.go")
		}
		// 更新 Video 数据结构中的 favorite_count 字段
		err = config.DB.Model(&model.Video{}).Where("id = ?", likeinfo.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			log.Println("UpdateVideoAfterUnLikeOperation favoritr_count error in ./service/favorite.go")
		}
	}
	return nil
}

func (like *FavoriteService) GetLikeList(listinfo model.FavoriteList) (model.FavoriteListResponse, error) {
	// 这个 favoritelist 中还有自己以及别人的视频
	// 首先需要判断用户是否处于登录状态，如果不是，需要登录
	// 登录用户的所有的点赞视频
	// 注意这个不仅需要获得自己的点赞列表，同样也需要获得别人的点赞列表
	var retList model.FavoriteListResponse
	// myId := listinfo.MyUserId
	// videolist, err := favoriteOperation.GetLikeList(listinfo.UserId)

	// if err != nil {
	// 	log.Println("GetLikeList error in ./service/favorite.go.")
	// 	return retList, err
	// }
	userId := listinfo.UserId
	var favorites []model.Favorite
	err := config.DB.Select("video_id").Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		log.Println("GetLikeList error in database/favorite.go")
		return retList, err
	}

	videoIDList := make([]uint, 0, len(favorites))
	log.Println("The favorite length is : ", len(favorites))
	for _, favorite := range favorites {
		videoIDList = append(videoIDList, favorite.VideoId)
	}

	// 根据 videoIDList 查找所有的 Video 数组
	var videoList []model.Video
	err = config.DB.Where("id in (?)", videoIDList).Find(&videoList).Error
	if err != nil {
		log.Println("GetVideoList error in database/video.go file")
		return retList, err
	}
	likevideolist, err := GetLikeVideoList(videoList)
	if err != nil {
		log.Println("GetLikeList error in ./service/favorite.go")
		return retList, err
	}
	retList.VideoList = likevideolist
	return retList, nil
}
