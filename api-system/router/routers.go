package router

import (
	"github.com/axiangcoding/antonstar-bot/controller/middleware"
	"github.com/axiangcoding/antonstar-bot/controller/v1"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/axiangcoding/antonstar-bot/swagger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	base := r.Group(settings.C().Server.BasePath)
	setRouterApiV1(base)
	return r
}

func setSwagger(r *gin.RouterGroup) {
	if settings.C().App.Swagger.Enable {
		swagger.SwaggerInfo.Version = settings.C().App.Version
		swagger.SwaggerInfo.Title = settings.C().App.Name
		swagger.SwaggerInfo.BasePath = r.BasePath()
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setRouterApiV1(r *gin.RouterGroup) {
	api := r.Group("/api")
	setSwagger(api)
	groupV1 := api.Group("/v1")
	{
		system := groupV1.Group("/system")
		{
			system.GET("/info", v1.SystemInfo)
		}
		cqhttp := groupV1.Group("/cqhttp")
		{
			cqhttpAuth := middleware.CqhttpAuth(
				settings.C().App.Service.CqHttp.SelfQQ,
				settings.C().App.Service.CqHttp.Secret)
			cqhttp.POST("/receive/event", cqhttpAuth, v1.CqHttpReceiveEvent)
			cqhttp.GET("/status", v1.CqHttpStatus)
		}
	}
}
