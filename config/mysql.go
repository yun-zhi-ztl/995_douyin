// Package config
// @author ufec https://github.com/ufec
// @date 2022/5/9
package config

// MysqlConfig
//  @Description: Mysql 配置
type MysqlConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Dbname   string `json:"db_name" mapstructure:"db_name"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Charset  string `json:"charset" yaml:"charset"`
	Prefix   string `json:"prefix" yaml:"prefix"`
}

// Dsn
//  @Description: 将数据库配置转为DSN格式链接
//  @receiver m *MysqlConfig
//  @return string	dsn
func (m *MysqlConfig) Dsn() string {
	// dsn格式: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}
