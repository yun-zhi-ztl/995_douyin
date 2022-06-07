/*
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-02 18:52:15
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 08:06:58
 * @FilePath: \GoPath\995_douyin\config\redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type RedisConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
