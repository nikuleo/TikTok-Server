package main

import (
	"TikTokServer/app/log"
)

func init() {
	log.InitLog()
}

func main() {
	defer log.Sync()
}
