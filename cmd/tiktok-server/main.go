package main

import (
	"TikTokServer/app/log"
	"TikTokServer/routes"
)

func init() {
	log.InitLog()
}

func main() {
	defer log.Sync()

	routes.Run()

	log.Info("Server exiting")
}
