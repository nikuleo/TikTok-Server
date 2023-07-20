package cache

import (
	"TikTokServer/pkg/config"
	"TikTokServer/pkg/tlog"
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	EXPIRE int = 60 // 60s 测试使用
)

var (
	Ctx              = context.Background()
	RdbUserInfo      *redis.Client
	RdbUserFavorite  *redis.Client
	RdbUserFollowing *redis.Client

	RdbVideoFarvite *redis.Client
	RdbVideoComment *redis.Client
)

// InitRedis 初始化Redis连接。
func InitRedis() {

	cfg := config.GetConfig("redisConfig").Viper
	address := cfg.GetString("redis.address")
	port := cfg.GetString("redis.port")
	password := cfg.GetString("redis.auth")

	RdbUserInfo = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       0, // 用户信息存入 DB0.
	})
	RdbUserFavorite = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       1, // 用户赞过的视频id存入 DB1.
	})
	RdbUserFollowing = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       2, // 关注列表信息信息存入 DB2.
	})

	RdbVideoFarvite = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       3, //  视频点赞信息存入 DB3.
	})
	RdbVideoComment = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       4, //  将视频评论信息存入 DB4.
	})

	tlog.Info("Redis init success.")
}
