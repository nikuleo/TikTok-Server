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
	RdbUserFollower  *redis.Client
	RdbUserFollowing *redis.Client
	RdbUserFriend    *redis.Client
	RdbVideoInfo     *redis.Client
	RdbVideoFarvite  *redis.Client
	RdbVideoComment  *redis.Client
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
	RdbUserFollower = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       1, // 粉丝列表信息存入 DB1.
	})
	RdbUserFollowing = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       2, // 关注列表信息信息存入 DB2.
	})
	RdbUserFriend = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       3, // 朋友列表信息存入 DB3.
	})

	RdbVideoInfo = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       4, // 视频信息存入 DB4.
	})
	RdbVideoFarvite = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       5, //  视频点赞信息存入 DB5.
	})
	RdbVideoComment = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       6, //  将视频评论信息存入 DB6.
	})

	tlog.Info("Redis init success.")
}
