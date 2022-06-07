/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 08:58:36
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 08:58:41
 * @FilePath: \GoPath\995_douyin\model\data\commen.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package data

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            uint   `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         uint   `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// 评论操作响应内容
type CommentResponse struct {
	BaseResponse Response `json:"response"`
	ID           uint     `json:"id"`          // 评论id
	Content      string   `json:"content"`     // 评论内容
	CreateDate   string   `json:"create_date"` // 评论发布日期，格式 mm-dd
	User         User     `json:"user"`        // 评论用户信息
}
