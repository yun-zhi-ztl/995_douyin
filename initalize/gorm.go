/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 09:14:54
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:23:25
 * @FilePath: \GoPath\995_douyin\initalize\gorm.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
		model.Comment{},
		model.Video{},
		model.Favorite{},
	)
	if err != nil {
		return err
	}
	return nil
}
