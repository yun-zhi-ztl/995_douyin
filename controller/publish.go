package controller

import (
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/service"
	"github.com/yun-zhi-ztl/995_douyin/utils"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var VIDEO_URL = "http://192.168.31.50:8080/public/video/"
var COVER_URL = "http://192.168.31.50:8080/public/cover/"
var VIDEO_PATH = "public/video"
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
		// currentUser := UsersLoginInfo[token]
		user_id, err := utils.ParserToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  "token解析失败",
			})
			return
		}
		currentUser, exist := service.QueryUser(user_id, user_id)
		if !exist {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  "用户查找失败",
			})
			return
		}
		coverName := strings.Split(file.Filename, ".")[0] + "_conver.png"
		// fmt.Println(path.Join(VIDEO_PATH, file.Filename))
		// fmt.Println(path.Join(COVER_PATH, coverName))
		err = utils.BuildThumbnailWithVideo(path.Join(VIDEO_PATH, file.Filename), path.Join(COVER_PATH, coverName))
		if err != nil {
			//若是生成封面失败，则使用默认的封面
			coverName = "123.png"
		}
		videoDB := model.Video{
			UserId:        uint(currentUser.Id),
			PlayUrl:       file.Filename,
			CoverUrl:      coverName,
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
	// user := UsersLoginInfo[token]
	token_id, err := utils.ParserToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 0,
			StatusMsg:  err.Error(),
		})
		return
	}
	user_str_id := c.Query("user_id")
	user_id, _ := strconv.Atoi(user_str_id)
	videoDB := model.Video{}
	var result *[]model.Video
	if user_id == 0 {
		result, err = videoDB.QueryByUserID(token_id)
	} else {
		result, err = videoDB.QueryByUserID(user_id)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 0,
			StatusMsg:  "查询失败",
		})
		return
	}
	videoListresponse := VideoListResponse{}
	response := Response{
		StatusCode: 0,
		StatusMsg:  "获取成功",
	}
	var videos = make([]Video, len(*result))
	for i := 0; i < len(*result); i++ {
		userinfo, exist := service.QueryUser(token_id, user_id)
		if !exist {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "something error"})
			return
		}
		videos[i].Author = User{
			Id:            uint(userinfo.Id),
			Name:          userinfo.Name,
			FollowCount:   int64(userinfo.FollowCount),
			FollowerCount: int64(userinfo.FollowerCount),
			IsFollow:      userinfo.IsFollow,
		}
		videos[i].Title = (*result)[i].Title
		videos[i].Id = (int64)((*result)[i].ID)
		videos[i].PlayUrl = VIDEO_URL + (*result)[i].PlayUrl
		videos[i].CoverUrl = COVER_URL + (*result)[i].CoverUrl
		videos[i].FavoriteCount = (int64)((*result)[i].FavoriteCount)
		videos[i].CommentCount = (int64)((*result)[i].CommentCount)
		videos[i].IsFavorite = IsFavorite(uint(token_id), (uint)((*result)[i].ID))
	}
	videoListresponse.Response = response
	videoListresponse.VideoList = videos
	c.JSON(http.StatusOK, videoListresponse)
}
