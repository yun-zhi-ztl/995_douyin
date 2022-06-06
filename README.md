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
## 运行
```shell
go run ./main.go ./router.go
```
### 编译
```shell
go build && ./simple-demo
```
### 代码结构
```text
├── config               # 配置文件
│   ├── config.go
│   ├── mysql.go
│   ├── redis.go
│   └── server.go
├── controller           # 控制器
│   ├── base.go
│   ├── comment.go
│   ├── common.go
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── initalize             # 初始化
│   ├── gormConfig
│   │   ├── gorm_config.go
│   ├── gorm_mysql.go
│   ├── gorm.go
│   └── viper.go
├── middleware
│   └── jwt.go
├── model
│   ├── data
│   │   ├── commen.go
│   │   └── demo_data.go
│   ├── comment.go
│   ├── favorite.go
│   ├── user.go
│   └── video.go
├── public
├── service
│   ├── comment.go
│   ├── favorite.go
│   ├── followService.go
│   ├── user.go
│   └── video.go
├── utils
│   ├── redis.go
│   ├── token.go
│   └── video.go
├── config.yaml
├── config.yaml.example
├── go.mod
├── go.sum
├── main.go
├── README.md
└── router.go
```

## 其他说明