package dao

import "douyin/config"

type VideoDB struct {
	VideoId       int    `json:"video_id,omitempty" gorm:"column:video_id"`
	UserId        int    `json:"user_id" gorm:"column:user_id"`
	PlayUrl       string `json:"play_url,omitempty" gorm:"column:play_url"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"column:cover_url"`
	Title         string `json:"title,omitempty" gorm:"column:title"`
	FavoriteCount int    `json:"favorite_count" gorm:"column:favorite_count"`
	CommentCount  int    `json:"comment_count" gorm:"column:comment_count"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"column:is_favotite"`
}

func (v VideoDB) TableName() string {
	return "douyin_video"
}

func (v VideoDB) Create() error {
	return config.DB.Create(&v).Error
}
