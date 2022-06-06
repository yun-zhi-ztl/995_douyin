/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-25 00:45:20
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-06 09:25:11
 * @FilePath: \GoPath\995_douyin\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/initalize"
	"github.com/yun-zhi-ztl/995_douyin/initalize/gormConfig"
)

func main() {
	initalize.InitViper()                           // 初始化Viper 读取配置文件
	config.GormConfig = gormConfig.InitGormConfig() // 初始化Gorm配置项
	config.DB = initalize.InitGorm()                // 初始化数据库
	if config.DB != nil {
		if err := initalize.CreateTable(config.DB); err != nil {
			panic(err)
		}
		println("数据库表初始化成功!")
	}
	r := gin.Default()
	initRouter(r)
	s := config.Config.Server.Port
	r.Run(":" + strconv.Itoa(s)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run(":8080")
}
