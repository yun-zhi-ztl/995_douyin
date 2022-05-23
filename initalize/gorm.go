// Package initalize
// @author ufec https://github.com/ufec
// @date 2022/5/9
package initalize

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"gorm.io/gorm"
)

// InitGorm
//  @Description: 初始化Gorm
//  @return *gorm.DB
func InitGorm() *gorm.DB {
	switch config.Config.Server.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// CreateTable
//  @Description: 根据模型自动创建表
func CreateTable(db *gorm.DB) error {
	// 创建表，自动迁移(把结构体和数据表进行对应)
	err := db.AutoMigrate(
		model.UserInfo{},
	//model.Video{},
	//model.Comment{},
	)
	if err != nil {
		return err
	}
	return nil
}
