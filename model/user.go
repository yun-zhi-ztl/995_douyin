package model

import "gorm.io/gorm"

// UserInfo 用户信息
type User struct {
	gorm.Model
	Name     string `gorm:"varchar(32);not null;unique"`
	Password string `gorm:"varchar(32);not null"`
}
