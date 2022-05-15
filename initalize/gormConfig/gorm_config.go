// Package gormConfig
// @author ufec https://github.com/ufec
// @date 2022/5/9
package gormConfig

import (
	"github.com/ufec/douyin_be/config"
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
