package routes

import (
	"TikTokServer/middleware"
	"TikTokServer/pkg/log"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	server *http.Server
)

func Run() {
	getRoutes()
	server = &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", log.Err(err))
		}
	}()

	// 优雅退出
	gracefulExit(server)
	// router.Run(":8080")
}

func getRoutes() {
	douyin := router.Group("/douyin", middleware.GinLog())
	addUserRoutes(douyin)
	addFeedRoutes(douyin)
	addPublishRoutes(douyin)
	addFavoriteRoutes(douyin)
	addCommentRoutes(douyin)
	addRelationRoutes(douyin)
	addMessageRoutes(douyin)
}

func gracefulExit(server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	sig := <-ch
	log.Info("receive exit signal", log.Any("", sig))
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		log.Error("server shutdown error", log.Err(err))
	}

	// 实际退出所耗费的时间
	log.Info("------exited--------", log.Duration("duration:", time.Since(now)))
}
