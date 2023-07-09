package main

import (
	"TikTokServer/model"
	"TikTokServer/pkg/auth"
	"TikTokServer/pkg/ossBucket"
	"TikTokServer/pkg/tlog"
	"TikTokServer/routes"
)

func Init() {
	tlog.InitLog()
	model.InitDB()
	auth.InitJWT()
	ossBucket.OssInit()
}

func main() {
	defer tlog.Sync()

	Init()

	routes.Run()

	tlog.Info("Server exiting")
}
