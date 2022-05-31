package controller

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/utils"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// token := c.PostForm("token")
	// title := c.PostForm("title")
	user_id := c.GetInt("id")

	// user_id, err := middleware.ParserToken(token)
	// if err != nil {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// 	})
	// 	return
	// }
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "file read false"},
		})
		return
	}
	nowTime := time.Now()
	saveDir := fmt.Sprintf("./public/video/%d/%d/%d/", nowTime.Year(), nowTime.Month(), nowTime.Day())
	if !utils.PathExists(saveDir) {
		if err := utils.MakeDir(saveDir); err != nil {
			// 创建目录失败
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "dir create false"},
			})
			return
		}
	}
	// 用户id_文件名_文件大小 拼接文件名 用于后续制作视频封面 提取文件后缀 后续对文件名进一步处理
	fileName, fileExt := fmt.Sprintf("%d_%s_%d", user_id, file.Filename, file.Size), filepath.Ext(file.Filename)
	// 用时间戳 对拼接后的文件名进行 hmac_sha256 散列 输出bas464格式
	saveFileName := utils.HmacSha256(fileName, strconv.FormatInt(nowTime.Unix(), 10), "base64")
	// 最终保存的目录+文件名组成为最终该文件被存储的路径
	saveVideoFile := filepath.Join(saveDir, saveFileName+fileExt)
	// 保存上传的视频文件
	if err := c.SaveUploadedFile(file, saveVideoFile); err != nil {
		// 保存视频失败
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "file save fasle"},
		})
		return
	}
	// 视频保存成功后 制作视频封面
	saveThumbnailFile := path.Join(saveDir, saveFileName+"_thumbnail.png")
	if err := utils.BuildThumbnailWithVideo(saveVideoFile, saveThumbnailFile); err != nil {
		// 封面生成失败
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "cover create false"},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "publis success"},
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

}
