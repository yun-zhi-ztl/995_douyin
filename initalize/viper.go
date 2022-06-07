/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-07 08:05:18
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:18:47
 * @FilePath: \GoPath\995_douyin\initalize\viper.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Package initalize
// @author ufec https://github.com/ufec
// @date 2022/5/9
package initalize

import (
	"errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yun-zhi-ztl/995_douyin/config"
)

// InitViper
//  @Description: 初始化Viper 用于解析 config.yaml / config.json 配置文件 解析后的配置统一通过 config.Config 获取
func InitViper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	viper.AddConfigPath("./") // call multiple times to add many search paths
	v.SetConfigType("yaml")
	readConfErr := v.ReadInConfig()
	if readConfErr != nil {
		panic(errors.New("打开配置文件出错，请检查根目录是否存在 config.yaml 文件 "))
	}
	if unmarshalConfigFile := v.Unmarshal(&config.Config); unmarshalConfigFile != nil {
		panic(errors.New("配置文件读取失败，请检查配置项与官方配置是否一致! "))
	}
	// 监听配置文件修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if unmarshalConfigFile := v.Unmarshal(&config.Config); unmarshalConfigFile != nil {
			println(errors.New("配置文件读取失败，请检查配置项与官方配置是否一致! "))
		}
	})
}
