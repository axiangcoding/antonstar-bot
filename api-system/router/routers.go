package router

import (
	"github.com/axiangcoding/antonstar-bot/controller/middleware"
	"github.com/axiangcoding/antonstar-bot/controller/v1"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/axiangcoding/antonstar-bot/swagger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// r.MaxMultipartMemory = 8 << 20
	// 全局中间件
	// 使用自定义中间件
	r.Use(middleware.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	setSession(r)
	setRouterV1(r)
	return r
}

func setSession(r *gin.Engine) {
	source := settings.Config.Data.Cache.Source
	address := strings.ReplaceAll(source, "redis://", "")
	address = strings.ReplaceAll(address, "/0", "")

	duration, err := time.ParseDuration(settings.Config.Auth.ExpireDuration)
	if err != nil {
		logging.Fatal(err)
	}

	store, err := redis.NewStore(1000, "tcp", address,
		"", []byte(settings.Config.Auth.Secret))
	if err != nil {
		logging.Fatal(err)
	}
	if err := redis.SetKeyPrefix(store, "Session:"); err != nil {
		logging.Fatal(err)
	}
	store.Options(sessions.Options{
		MaxAge:   int(duration.Seconds()),
		Path:     "-",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode})
	r.Use(sessions.Sessions("session", store))
}

func setSwagger(r *gin.RouterGroup) {
	if settings.Config.App.Swagger.Enable {
		swagger.SwaggerInfo.Version = settings.Config.App.Version
		swagger.SwaggerInfo.Title = settings.Config.App.Name
		swagger.SwaggerInfo.BasePath = settings.Config.Server.BasePath
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setRouterV1(r *gin.Engine) {
	base := r.Group(settings.Config.Server.BasePath)
	setSwagger(base)
	groupV1 := base.Group("/v1")
	{
		system := groupV1.Group("/system")
		{
			system.GET("/info", v1.SystemInfo)
		}
		cqhttp := groupV1.Group("/cqhttp")
		{
			cqhttpAuth := middleware.CqhttpAuth(
				settings.Config.Service.CqHttp.SelfQQ,
				settings.Config.Service.CqHttp.Secret)
			cqhttp.POST("/receive/event", cqhttpAuth, v1.CqHttpReceiveEvent)
			cqhttp.GET("/status", v1.CqHttpStatus)
		}
	}
}
