package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yun-zhi-ztl/995_douyin/config"
	"github.com/yun-zhi-ztl/995_douyin/model"
	"github.com/yun-zhi-ztl/995_douyin/utils"

	"github.com/garyburd/redigo/redis"
)

func Follow(userId int, targetId int) {
	userFolloweeRedisKey := utils.GetFolloweeRedisKey(userId)
	userFollowerRedisKey := utils.GetFollowerRedisKey(targetId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	// 执行一个有序zset插入
	_, err = conn.Do("ZADD", userFolloweeRedisKey, time.Now().Unix(), strconv.Itoa(targetId))
	if err != nil {
		fmt.Println("redis set followee failed:", err)
	}
	// 再执行一个有序zset插入
	_, err = conn.Do("ZADD", userFollowerRedisKey, time.Now().Unix(), strconv.Itoa(userId))
	if err != nil {
		fmt.Println("redis set follower failed:", err)
	}
}

func Unfollow(userId int, targetId int) {
	userFolloweeRedisKey := utils.GetFolloweeRedisKey(userId)
	userFollowerRedisKey := utils.GetFollowerRedisKey(targetId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	// 执行一个有序zset删除
	_, err = conn.Do("Zrem", userFolloweeRedisKey, strconv.Itoa(targetId))
	if err != nil {
		fmt.Println("redis rem followee failed:", err)
	}
	// 再执行一个有序zset删除
	_, err = conn.Do("Zrem", userFollowerRedisKey, strconv.Itoa(userId))
	if err != nil {
		fmt.Println("redis rem follower failed:", err)
	}
}

// HasFollow 查询当前用户是否已关注实体
func HasFollow(userId int, targetId int) bool {
	userFolloweeRedisKey := utils.GetFolloweeRedisKey(userId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()

	res, _ := conn.Do("zscore", userFolloweeRedisKey, strconv.Itoa(targetId))
	if res != nil {
		return true
	} else {
		return false
	}
}

// FolloweeCount 查询关注的人的数量
func FolloweeCount(userId int) int64 {
	userFolloweeRedisKey := utils.GetFolloweeRedisKey(userId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	num, err := conn.Do("zcard", userFolloweeRedisKey)
	if err != nil {
		return 0
	} else {
		return num.(int64)
	}
}

// FollowerCount 查询粉丝的数量
func FollowerCount(userId int) int64 {
	userFollowerRedisKey := utils.GetFollowerRedisKey(userId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	num, err := conn.Do("zcard", userFollowerRedisKey)
	if err != nil {
		return 0
	} else {
		return num.(int64)
	}
}

// GetRedisConn 把建立redis数据库连接方法抽出来
func GetRedisConn() redis.Conn {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	return conn
}

type Userinfo struct {
	Id   int
	Name string
}

// FindFollowees 查询关注的人,返回[]Userinfo
func FindFollowees(userId int) []Userinfo {
	userFolloweeRedisKey := utils.GetFolloweeRedisKey(userId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	users := make([]Userinfo, 0, FolloweeCount(userId))
	//var users []Userinfo
	res, _ := redis.Strings(conn.Do("zrange", userFolloweeRedisKey, 0, -1))
	fmt.Println("res len:", len(res))
	if err != nil {
		fmt.Println("zrange failed", err.Error())
	} else {
		for _, v := range res {
			id, _ := strconv.Atoi(v)
			var user model.UserInfo
			config.DB.Where("Id = ?", id).Find(&user)
			users = append(users, Userinfo{int(user.ID), user.UserName})
		}
	}
	return users
}

// FindFollowers 查询粉丝,返回[]Userinfo
func FindFollowers(userId int) []Userinfo {
	userFollowerRedisKey := utils.GetFollowerRedisKey(userId)
	// 连接redis数据库,指定数据库的IP和端口
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	} else {
		fmt.Println("Connect to redis ok.")
	}
	// 函数退出时关闭连接
	defer conn.Close()
	users := make([]Userinfo, 0, FollowerCount(userId))
	//var users []Userinfo
	res, err := redis.Strings(conn.Do("zrange", userFollowerRedisKey, 0, -1))
	if err != nil {
		fmt.Println("zrange failed", err.Error())
	} else {
		for _, v := range res {
			id, _ := strconv.Atoi(v)
			var user model.UserInfo
			config.DB.Where("Id = ?", id).Find(&user)
			users = append(users, Userinfo{int(user.ID), user.UserName})
		}
	}
	return users
}
