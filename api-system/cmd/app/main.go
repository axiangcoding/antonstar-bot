package main

import (
	"context"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/cache"
	"github.com/axiangcoding/antonstar-bot/controller/validation"
	"github.com/axiangcoding/antonstar-bot/cron"
	"github.com/axiangcoding/antonstar-bot/data"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/router"
	"github.com/axiangcoding/antonstar-bot/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	logging.Setup()
	data.Setup()
	cache.Setup()
	cron.Setup()
	validation.Setup()
}

// @title        axiangcoding/anton-star
// @version      1.0.0
// @description  api system build by ax-web
// @termsOfService

// @contact.name  axiangcoding
// @contact.url
// @contact.email  axiangcoding@gmail.com

// @license.name
// @license.url

// @accept   json
// @produce  json

// @securityDefinitions.apikey  AppToken
// @in                          query
// @name                        app_token
func main() {
	runMode := settings.Config.Server.RunMode
	gin.SetMode(runMode)
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", settings.Config.Server.Port),
		Handler: r,
	}
	// Initialize the server in the goroutine so that it will not block the graceful stop processing below
	// 在goroutine中初始化服务器，这样就不会阻塞下文的优雅停止处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal("Server error. ", err)
		}
	}()
	// Wait for the interrupt signal to gracefully stop the server, set a delay of 5 seconds
	// 等待中断信号来优雅停止服务器，设置的5秒延迟
	quit := make(chan os.Signal, 1)
	// kill 	syscall.SIGTERM
	// kill -2 	syscall.SIGINT
	// kill -9 	syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Info("Shutting down server...")
	// ctx is used to notify the server that there is 5 seconds left to end the request currently being processed
	// ctx是用来通知服务器还有5秒的时间来结束当前正在处理的request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Fatal("Server forced to shutdown. ", err)
	}

	logging.Info("Server exiting")
}
