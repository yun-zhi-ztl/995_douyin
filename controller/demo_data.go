package controller

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            4,
	Name:          "04",
	FollowCount:   1,
	FollowerCount: 0,
	IsFollow:      false,
}

var DemoVideo = Video{
	Id:            1,
	Author:        DemoUser,
	PlayUrl:       "http://192.168.0.6:8081/public/fa3c3f32eef6952f1f9e01743cd281a7.mp4",
	CoverUrl:      "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farticle%2F180a319f682220a5fe8c6cdd2408da51efb0223e.jpg&refer=http%3A%2F%2Fi0.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1656092757&t=75537e6918e823c3cb16929d02de1866",
	FavoriteCount: 100,
	CommentCount:  90,
	IsFavorite:    false,
}
