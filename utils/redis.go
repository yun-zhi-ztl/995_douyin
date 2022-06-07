/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 18:51:00
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-02 18:51:08
 * @FilePath: \GoPath\995_douyin\utils\redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import "strconv"

func GetFolloweeRedisKey(userId int) string {
	return "followee" + ":" + strconv.Itoa(userId)
}

func GetFollowerRedisKey(userId int) string {
	return "follower" + ":" + strconv.Itoa(userId)
}
