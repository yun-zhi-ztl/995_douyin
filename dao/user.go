package dao

import "douyin/config"

type UserDB struct {
	ID         int       `json:"id"`
	UserName   string    `json:"usrName"`
	PassWord   int       `json:"password"`
	Token      string    `json:token`
	CreateTime time.Time `json:"createTime"`
}

func (UserDB) TableName() string {
	return "douyin_user"
}

func (u UserDB) Create() error {
	return config.DB.Create(&u).Error
}
