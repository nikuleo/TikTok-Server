package main

import (
	"TikTokServer/model"
	"TikTokServer/pkg/tlog"
	"TikTokServer/routes"
)

func Init() {
	tlog.InitLog()
	model.InitDB()
}

func main() {
	defer tlog.Sync()

	Init()

	routes.Run()

	tlog.Info("Server exiting")
}
