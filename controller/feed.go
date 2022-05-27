package controller

import (
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

}
