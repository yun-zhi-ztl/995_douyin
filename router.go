/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-07 08:05:18
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:09:51
 * @FilePath: \GoPath\995_douyin\router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"github.com/yun-zhi-ztl/995_douyin/controller"
	"github.com/yun-zhi-ztl/995_douyin/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.Use(gin.Recovery())
	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.JWTAuth("query"), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/publish/list/", middleware.JWTAuth("query"), controller.PublishList)
	apiRouter.POST("/publish/action/", middleware.JWTAuth("form"), controller.Publish)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
