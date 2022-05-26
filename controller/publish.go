package controller

import (
	"douyin/dao"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var VIDEO_URL = "http:192.168.1.118:8094/public/video"
var VIDEO_PATH = "public/video"
var COVER_URL = "http:192.168.1.118:8094/public/cover"
var COVER_PATH = "public/cover"

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  "文件上传失败",
		})
	} else {
		c.SaveUploadedFile(file, path.Join(VIDEO_PATH, file.Filename))
		title := c.PostForm("title")
		token := c.PostForm("token")
		currentUser := UsersLoginInfo[token]
		videoDB := dao.VideoDB{
			UserId:        int(currentUser.Id),
			PlayUrl:       path.Join(VIDEO_URL, file.Filename),
			CoverUrl:      path.Join(COVER_URL, "123.png"),
			Title:         title,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		}
		err := videoDB.Create()
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  "文件上传失败",
			})
		} else {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 0,
				StatusMsg:  "文件成功上传",
			})
		}

	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

}
