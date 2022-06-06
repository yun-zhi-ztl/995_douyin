# 995队-代码敲不队

## 人员分工：
- 朱太龙：队长、用户注册与登录、评论操作与评论列表；
- 许意：框架设计、MySQL数据库设计、用户信息接口、视频流接口；
- 周冰：MySQL数据库设计、点赞操作、点赞列表；
- 周彪：Redis数据库相关、关注列表、粉丝列表、关系操作接口；
- 刘建军：MySQL数据库设计、视频投稿、发布列表接口；


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