/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 08:58:49
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 09:59:35
 * @FilePath: \GoPath\995_douyin\model\data\demo_data.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package data

var DemoVideos = []Video{
	{
		Id:            42,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.31.50:8080/public/test.mp4",
		CoverUrl:      "http://192.168.31.50:8080/public/feng.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser1,
		PlayUrl:       "http://192.168.31.50:8080/public/test1.mp4",
		CoverUrl:      "http://192.168.31.50:8080/public/feng.png",
		FavoriteCount: 1008,
		CommentCount:  20,
		IsFavorite:    false,
	},
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.31.50:8080/public/test2.mp4",
		CoverUrl:      "http://192.168.31.50:8080/public/feng.png",
		FavoriteCount: 108,
		CommentCount:  20,
		IsFavorite:    false,
	},
}
var DemoUser = User{
	Id:   1,
	Name: "TestUser",
}
var DemoUser1 = User{
	Id:   2,
	Name: "TestUser",
}
var DemoUser2 = User{
	Id:   3,
	Name: "TestUser",
}
