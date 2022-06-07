/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-07 08:05:18
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:19:32
 * @FilePath: \GoPath\995_douyin\initalize\gormConfig\gorm_config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Package gormConfig
// @author ufec https://github.com/ufec
// @date 2022/5/9
package gormConfig

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitGormConfig
//  @Description: 初始化Gorm的部分配置项
//  @return *gorm.Config
func InitGormConfig() *gorm.Config {
	var tablePrefix string
	switch config.Config.Server.DbType {
	case "mysql":
		tablePrefix = config.Config.Mysql.Prefix
	default:
		tablePrefix = config.Config.Mysql.Prefix
	}
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		PrepareStmt: true,
		//FullSaveAssociations:                     false,
		//Logger:                                   nil,
		//NowFunc:                                  nil,
		//DryRun:                                   false,
		//DisableAutomaticPing:                     false,
		//DisableForeignKeyConstraintWhenMigrating: false,
		//DisableNestedTransaction:                 false,
		//AllowGlobalUpdate:                        false,
		//QueryFields:                              false,
		//CreateBatchSize:                          0,
		//ClauseBuilders:                           nil,
		//ConnPool:                                 nil,
		//Dialector:                                nil,
		//Plugins:                                  nil,
	}
	return gormConfig
}
