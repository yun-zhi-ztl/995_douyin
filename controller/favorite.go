/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-05-15 22:11:28
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:00:09
 * @FilePath: \GoPath\995_douyin\controller\favorite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yun-zhi-ztl/995_douyin/middleware"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/service"
)

var favoriteService service.FavoriteService

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var favoriteaction model.FavoriteAction
	err := c.ShouldBindQuery(&favoriteaction)
	if err != nil {
		model.Failed("bind error in controller/favorite.go", c)
		return
	}
	user := model.UserInfo{}
	token := c.Query("token")
	user.ID, err = middleware.GetUidByToken(token)
	if err != nil {
		log.Println("Get uid by token error in ./database/video.go")
	}

	favoriteaction.UserId = user.ID
	if favoriteaction.ActionType == 1 {
		log.Println("点赞操作")
	} else {
		log.Println("取消点赞操作")
	}

	err = favoriteService.LikeOperation(favoriteaction)
	if err != nil {
		// log.Println("error in FavoriteAction in ./controller/favorite.go")
		model.Failed("error in FavoriteAction in ./controller/favorite.go", c)
		return
	}
	model.LikeOperationSuccess("Like operation success", c)

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var favoritelist model.FavoriteList

	token := c.Query("token")

	favoritelist.UserId, _ = middleware.GetUidByToken(token)
	favoritelist.Token = token
	log.Println("token : ", token)
	log.Println("Favoritelist userid is : ", favoritelist.UserId)
	retlist, err := favoriteService.GetLikeList(favoritelist)
	if err != nil {
		// log.Println("error in FavoriteAction in ./controller/favorite.go")
		model.Failed("error in FavoriteList in ./controller/favorite.go", c)
		return
	}
	model.GetLikeListSuccess("Get like list suucess", c, retlist.VideoList)
}
