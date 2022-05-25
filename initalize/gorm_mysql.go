// Package initalize
// @author ufec https://github.com/ufec
// @date 2022/5/9
package initalize

import (
	"douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql
//  @Description: 初始化Gorm with Mysql
//  @return *gorm.DB
func GormMysql() *gorm.DB {
	m := config.Config.Mysql
	mysqlConfig := mysql.Config{
		DriverName:                "",
		ServerVersion:             "",
		DSN:                       m.Dsn(),
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}
	db, openGormErr := gorm.Open(mysql.New(mysqlConfig), config.GormConfig)
	if openGormErr != nil {
		panic(openGormErr)
		return nil
	}
	return db
}
