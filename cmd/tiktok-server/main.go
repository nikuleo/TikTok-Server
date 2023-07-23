package main

import (
	"TikTokServer/cache"
	"TikTokServer/model"
	"TikTokServer/pkg/auth"
	"TikTokServer/pkg/ossBucket"
	"TikTokServer/pkg/tlog"
	"TikTokServer/rabbitmq"
	"TikTokServer/routes"
)

func Init() {
	tlog.InitLog()
	model.InitDB()
	auth.InitJWT()
	ossBucket.OssInit()
	cache.InitRedis()
	rabbitmq.InitRabbitMQ()
}

func main() {
	defer tlog.Sync()
	defer rabbitmq.MQClose()

	Init()

	routes.Run()

	tlog.Info("Server exiting")
}
