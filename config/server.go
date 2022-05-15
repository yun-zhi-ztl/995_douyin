// Package config
// @author ufec https://github.com/ufec
// @date 2022/5/9
package config

// ServerConfig
//  @Description: 系统配置
type ServerConfig struct {
	Port   int    `json:"port,omitempty" yaml:"port"`
	DbType string `json:"db_type,omitempty" yaml:"db_type"`
}
