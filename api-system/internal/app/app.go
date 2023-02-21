package app

import (
	"context"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/internal/cache"
	"github.com/axiangcoding/antonstar-bot/internal/controller/http/v1"
	"github.com/axiangcoding/antonstar-bot/internal/cron"
	"github.com/axiangcoding/antonstar-bot/internal/data"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initProject() {
	setting.InitConf()
	logging.InitLogger()
	data.InitData()
	cache.Setup()
	cron.Setup()
}

func Run() {
	initProject()
	runMode := setting.C().Server.RunMode
	gin.SetMode(runMode)
	r := v1.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.C().Server.Port),
		Handler: r,
	}
	// Initialize the server in the goroutine so that it will not block the graceful stop processing below
	// 在goroutine中初始化服务器，这样就不会阻塞下文的优雅停止处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.L().Fatal("Server error. ", logging.Error(err))
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
	logging.L().Info("Shutting down server...")
	// ctx is used to notify the server that there is 5 seconds left to end the request currently being processed
	// ctx是用来通知服务器还有5秒的时间来结束当前正在处理的request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.L().Fatal("Server forced to shutdown. ", logging.Error(err))
	}

	logging.L().Info("Server exiting")
}
