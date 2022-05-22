package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/controller"
	"github.com/yun-zhi-ztl/995_douyin/middleware"
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
	// apiRouter.POST("/user/token/", controller.Token)
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
