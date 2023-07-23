<div align="center">

  <h1 align="center">TikTok-Server</h1>
  <span>极简抖音后端服务 —— tiny tiktok server</span></br>
  <span>字节跳动青训营-后端结营项目</span></br></br>

  <span>![license](https://img.shields.io/badge/license-MIT-green)</span> <span>![language](https://img.shields.io/badge/language-Go-blue)</span> <span>![青训营](https://img.shields.io/badge/%E5%AD%97%E8%8A%82%E9%9D%92%E8%AE%AD%E8%90%A5-%E7%AC%AC%E4%B8%89%E5%B1%8A-red)</span> <span>![score](https://img.shields.io/badge/%E8%8E%B7%E5%A5%96-%E8%B6%85%E7%BA%A7%E7%A0%81%E5%8A%9B-pink)</span>

</div>

## 介绍
本仓库是我的字节跳动青训营第三届结营项目个人实现仓库（以个人的思路与设计风格实现，做到代码简洁规范）。  
原小队合作仓库：[八位帝皇丸小队仓库](https://github.com/jhzol/TikTokLite)   

若有问题或建议欢迎 issue 讨论，十分感谢 :smiley_cat:

### 任务列表  
**已完成：**
- [x] 代码规范，commit 清晰
- [x] 项目架构设计，各层职责分明
- [x] 关系数据库设计规范，字段类型正确，存储大小设计合理，索引正确
- [x] 鉴权、日志、recovery 的 Gin 中间件
- [x] 错误码、配置文件、日志封装
- [x] 视频文件上传 oss 对象存储
- [x] MySQL 多表更新使用事务实现
- [x] 完成各 api 接口业务逻辑
- [x] postman 各接口测试 

**正在做：**
- [ ] 实现 redis 缓存服务 
- [ ] 修改读业务先查缓存，再查数据库
- [ ] 增加 rabbitMQ 消息队列服务
- [ ] 整体项目 docker 部署

**待定：**
- [ ] 完善单元测试
- [ ] 使用 pprof 性能测试
- [ ] 重构为微服务架构
---------
## Clean MVC 架构
### 项目布局
```bash
.
├── cmd
│   └── tiktok-server           # 项目启动入口
├── .config                     # 服务相关 yaml 配置文件
├── routes                      # 路由层
├── controller                  # 控制层
├── service                     # 服务层
├── cache                       # redis 缓存层
├── model                       # 持久化层
├── middleware                  # gin 中间件
├── pkg                         # 公共包
│   ├── auth                    # JWT 鉴权
│   ├── config                  # 封装 viper 配置包
│   ├── errorcode               # 错误码
│   ├── ossBucket               # 阿里云 oss 对象存储
│   ├── response                # 统一响应包装
│   ├── tlog                    # 日志
│   └── util                    # 工具包
├── docker                      # docker-compose 配置
│   ├── mysql
│   └── redis
│       ├── data                # redis 持久化数据（aof 与 快照）
│       │   └── appendonlydir
│       └── log
└── idl                         # protobuf idl 文件与生成的 go 文件
    ├── gen
    └── proto
```
## redis 缓存配置
- 最大内存限制 100mb （电脑内存小，docker 启动 redis 内存占用太高）
- 开启超出内存调度策略 LFU
- 启动主动整理内存碎片
- 使用混合持久化（AOF + RDB）

### 缓存与数据库一致性方案
读写分离，写操作先更新数据库，再删除缓存。其中删除缓存操作使用消息队列重试缓存删除。


## 日志方案
- zap
- lumberjack

## 路由
- gin

## 错误代码封装
> 使用该仓库的错误代码方案
> https://github.com/marmotedu/sample-code



## 微服务
// TODO: 之后有时间做(框架待定：kratos)

## 鸣谢
八位帝皇丸小队成员  
[字节跳动青训营](https://youthcamp.bytedance.com/)