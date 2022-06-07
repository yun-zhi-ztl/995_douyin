// Package config
// @author ufec https://github.com/ufec
// @date 2022/5/9
package config

import "gorm.io/gorm"

// Conf
//  @Description: 系统所有配置项
type Conf struct {
	Mysql  *MysqlConfig
	Server *ServerConfig
	Redis  *RedisConfig
}

var (
	Config       *Conf
	GormConfig   *gorm.Config
	DB           *gorm.DB
	ServerDomain = "http://192.168.31.50:8080/" // 用于访问视频和图片
)
