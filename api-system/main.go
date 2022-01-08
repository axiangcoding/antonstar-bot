package main

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/pkg/auth"
	"axiangcoding/antonstar/api-system/pkg/cache"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"axiangcoding/antonstar/api-system/pkg/mq"
	"axiangcoding/antonstar/api-system/pkg/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
	logging.Setup()
	data.Setup()
	cache.Setup()
	mq.Setup()
	auth.Setup()
}

// @title        Golang Gin Template API
// @version      1.0.0
// @description  An example of gin
// @termsOfService

// @contact.name  axiangcoding
// @contact.url
// @contact.email  axiangcoding@gmail.com

// @license.name
// @license.url

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        token

// @accept   json
// @produce  json
func main() {
	runMode := conf.Config.Server.RunMode
	gin.SetMode(runMode)
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Config.Server.Port),
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
