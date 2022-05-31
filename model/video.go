package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"int;not null"`
	PlayUrl  string `gorm:"varchar(32);not null"`
	CoverUrl string `gorm:"varchar(32);not null"`
}
