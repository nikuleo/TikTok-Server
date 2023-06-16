package main

import (
	"TikTokServer/model"
	"TikTokServer/pkg/log"
	"TikTokServer/routes"
)

func Init() {
	log.InitLog()
	model.InitDB()
}

func main() {
	defer log.Sync()

	Init()

	routes.Run()

	log.Info("Server exiting")
}
