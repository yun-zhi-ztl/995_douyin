package controller

import (
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//userId := Qualify(c).ID
	//获取文件

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

}
