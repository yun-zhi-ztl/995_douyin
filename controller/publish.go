package controller

import (
	"douyin/dao"
	"douyin/utils"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

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
		coverName := strings.Split(file.Filename, ".")[0] + "_conver.png"
		fmt.Println(path.Join(VIDEO_PATH, file.Filename))
		fmt.Println(path.Join(COVER_PATH, coverName))
		err := utils.BuildThumbnailWithVideo(path.Join(VIDEO_PATH, file.Filename), path.Join(COVER_PATH, coverName))
		if err != nil {
			//若是生成封面失败，则使用默认的封面
			coverName = "123.png"
		}
		videoDB := dao.VideoDB{
			UserId:        int(currentUser.Id),
			PlayUrl:       path.Join(VIDEO_URL, file.Filename),
			CoverUrl:      path.Join(COVER_URL, coverName),
			Title:         title,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		}
		err = videoDB.Create()
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  "文件上传失败",
			})
		} else {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "文件成功上传",
			})
		}

	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user := UsersLoginInfo[token]
	user_str_id := c.Query("user_id")
	user_id, _ := strconv.Atoi(user_str_id)
	videoDB := dao.VideoDB{}
	result, err := videoDB.QueryByUserID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 0,
			StatusMsg:  "查询失败",
		})
	} else {
		videoListresponse := VideoListResponse{}
		response := Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		}
		var videos = make([]Video, len(*result), len(*result))
		for i := 0; i < len(*result); i++ {
			videos[i].Author = user
			videos[i].Title = (*result)[i].Title
			videos[i].Id = (int64)((*result)[i].VideoId)
			videos[i].PlayUrl = (*result)[i].PlayUrl
			videos[i].CoverUrl = (*result)[i].CoverUrl
			videos[i].FavoriteCount = (int64)((*result)[i].FavoriteCount)
			videos[i].CommentCount = (int64)((*result)[i].CommentCount)
			videos[i].IsFavorite = (*result)[i].IsFavorite
		}
		videoListresponse.Response = response
		videoListresponse.VideoList = videos
		c.JSON(http.StatusOK, videoListresponse)
	}

}
