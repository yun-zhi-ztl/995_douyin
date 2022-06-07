package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/yun-zhi-ztl/995_douyin/config"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/service"
	"github.com/yun-zhi-ztl/995_douyin/utils"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	userId := c.GetInt("userID")
	if userId == 0 {
		Failed(c, "用户不存在")
		return
	}
	file, getUploadFileErr := c.FormFile("data")
	if getUploadFileErr != nil {
		Failed(c, getUploadFileErr.Error())
		return
	}
	pwd, getPwdErr := os.Getwd()
	if getPwdErr != nil {
		Failed(c, getPwdErr.Error())
		return
	}
	nowTime, saveDir := time.Now(), ""
	if runtime.GOOS == "windows" {
		saveDir = fmt.Sprintf(".\\public\\%d\\%d\\%d\\", nowTime.Year(), nowTime.Month(), nowTime.Day())
	} else {
		saveDir = fmt.Sprintf("./public/%d/%d/%d/", nowTime.Year(), nowTime.Month(), nowTime.Day())
	}
	// 不存在该目录则自动创建
	if !utils.PathExists(saveDir) {
		if mkDirErr := utils.MakeDir(saveDir); mkDirErr != nil {
			Failed(c, mkDirErr.Error())
			return
		}
	}
	// 用户id_文件名_文件大小 拼接文件名 用于后续制作视频封面 提取文件后缀 后续对文件名进一步处理
	fileName := fmt.Sprintf("%d_%d_%s", userId, file.Size, file.Filename)
	fileExt := filepath.Ext(fileName)
	// 用时间戳 对拼接后的文件名进行 hmac_sha256 散列 输出bas464格式
	saveFileName := utils.HmacSha256(fileName, strconv.FormatInt(nowTime.Unix(), 10), "hex")
	// 最终保存的目录+文件名组成为最终该文件被存储的路径
	saveVideoFile := filepath.Join(saveDir, saveFileName+fileExt)
	// 保存上传的视频文件
	if err := c.SaveUploadedFile(file, filepath.Join(pwd, saveVideoFile)); err != nil {
		Failed(c, "保存视频文件失败")
		return
	}
	// 视频保存成功后 制作视频封面
	saveThumbnailFile := filepath.Join(saveDir, saveFileName+"_thumbnail.png")
	if err := utils.BuildThumbnailWithVideo(filepath.Join(pwd, saveVideoFile), filepath.Join(pwd, saveThumbnailFile)); err != nil {
		Failed(c, "封面图生成失败")
		return
	}
	title := c.PostForm("title")
	if _, createErr := videoService.Create(saveVideoFile, saveThumbnailFile, title, uint(userId)); createErr != nil {
		Failed(c, createErr.Error())
		return
	}
	Success(c, "上传成功")
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
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "user_id不存在"})
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
		videos[i].PlayUrl = config.ServerDomain + (*result)[i].PlayUrl
		videos[i].CoverUrl = config.ServerDomain + (*result)[i].CoverUrl
		videos[i].FavoriteCount = (int64)((*result)[i].FavoriteCount)
		videos[i].CommentCount = (int64)((*result)[i].CommentCount)
		videos[i].IsFavorite = IsFavorite((uint)((*result)[i].ID), uint(token_id))
	}
	videoListresponse.Response = response
	videoListresponse.VideoList = videos
	c.JSON(http.StatusOK, videoListresponse)
}
