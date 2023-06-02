## TIKTOK-SERVER

## 架构
### MVC 结构
#### 项目 layout
```bash
.
├── app
│   └── log
│       └── log.go
├── cmd
│   └── tiktok-server
│       └── main.go
├── config
├── controller
├── go.mod
├── go.sum
├── idl
│   ├── gen
│   │   ├── comment.pb.go
│   │   ├── favorite.pb.go
│   │   ├── feed.pb.go
│   │   ├── message.pb.go
│   │   ├── publish.pb.go
│   │   ├── relation.pb.go
│   │   └── user.pb.go
│   └── proto
│       ├── comment.proto
│       ├── favorite.proto
│       ├── feed.proto
│       ├── message.proto
│       ├── publish.proto
│       ├── relation.proto
│       └── user.proto
├── middleware
├── model
├── README.md
└── routes
```
### 微服务
// TODO: 待定


## 日志方案
- zap
- lumberjack

## 路由
- gin