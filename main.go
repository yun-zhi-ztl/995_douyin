package main

import (
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
	// s := config.Config.Server.Port
	// r.Run(":" + strconv.Itoa(s)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8080")
}
