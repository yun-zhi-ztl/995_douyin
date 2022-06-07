/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-07 08:05:36
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:12:49
 * @FilePath: \GoPath\995_douyin\initalize\gorm_mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Package initalize
// @author ufec https://github.com/ufec
// @date 2022/5/9
package initalize

import (
	"github.com/yun-zhi-ztl/995_douyin/config"
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
	}
	return db
}
