<!--
 * @Author: yun-zhi-ztl 15071461069@163.com
 * @Date: 2022-06-07 08:05:36
 * @LastEditors: yun-zhi-ztl 15071461069@163.com
 * @LastEditTime: 2022-06-07 09:12:19
 * @FilePath: \GoPath\995_douyin\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# 995队-代码敲不队

## 项目运行
### MySQL安装
### Redis安装
### 配置修改
- config/config.go
```go
var (
	Config       *Conf
	GormConfig   *gorm.Config
	DB           *gorm.DB
    // 此处需要该为服务器的ip和端口号
	ServerDomain = "http://192.168.31.50:8080/" // 用于访问视频和图片
)
```
- config.yaml
```yaml
server:
  port: 8080            # 服务端口号
  db_type: mysql
  host: 127.0.0.1

mysql:
  host: 127.0.0.1
  port: 3306
  username: root            # mysql用户名
  password: winner0423      # mysql密码
  charset: utf8mb4
  prefix: douyin_
  db_name: douyin           # 提前在mysql内新建一个douyin的database
  
redis:                      # 提前开启redis
  host: 127.0.0.1
  port: 6379
```
## 运行
```shell
go run ./main.go ./router.go
```
### 编译
```shell
go build && ./simple-demo
```
## 代码结构
```text
├── config                  # 配置文件
│   ├── config.go
│   ├── mysql.go
│   ├── redis.go
│   └── server.go
├── controller              # 控制器
│   ├── base.go
│   ├── comment.go
│   ├── common.go           # 共有信息
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── initalize               # 初始化
│   ├── gormConfig
│   │   ├── gorm_config.go
│   ├── gorm_mysql.go
│   ├── gorm.go
│   └── viper.go
├── middleware              # 中间件
│   └── jwt.go
├── model                   # 模型
│   ├── data
│   │   ├── commen.go
│   │   └── demo_data.go
│   ├── comment.go
│   ├── favorite.go
│   ├── user.go
│   └── video.go
├── public                  # 视频、封面保存位置
├── service                 # 逻辑层部分
│   ├── comment.go
│   ├── favorite.go
│   ├── followService.go
│   ├── user.go
│   └── video.go
├── utils                   # 工具类
│   ├── redis.go
│   ├── token.go
│   └── video.go
├── config.yaml             # 配置文件
├── config.yaml.example
├── go.mod
├── go.sum
├── main.go
├── README.md
└── router.go
```

## 其他说明