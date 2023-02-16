package router

import (
	"github.com/axiangcoding/antonstar-bot/controller"
	"github.com/axiangcoding/antonstar-bot/controller/middleware"
	"github.com/axiangcoding/antonstar-bot/controller/v1"
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/axiangcoding/antonstar-bot/static/swagger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// r.MaxMultipartMemory = 8 << 20
	// 全局中间件
	logger := logging.L()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.CustomRecoveryWithZap(logger, true, func(c *gin.Context, err any) {
		app.ServerFailed(c, http.StatusInternalServerError)
	}))
	base := r.Group(settings.C().Server.BasePath)
	setWebResources(base)
	setRouterApiV1(base)
	return r
}

func setWebResources(r *gin.RouterGroup) {
	r.GET("/", controller.BaseRedirect)
	r.StaticFS("/web", http.Dir("web"))
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
		app := groupV1.Group("/app")
		{
			app.GET("/info", v1.AppInfo)
		}
		cqhttp := groupV1.Group("/cqhttp")
		{
			cqhttpAuth := middleware.CqhttpAuth(
				settings.C().App.Service.CqHttp.SelfQQ,
				settings.C().App.Service.CqHttp.Secret)
			cqhttp.POST("/receive/event", cqhttpAuth, v1.CqHttpReceiveEvent)
			cqhttp.GET("/status", v1.CqHttpStatus)
		}
		wt := groupV1.Group("/wt")
		{
			wt.GET("/profile", v1.GameUserProfile)
		}
	}
}

func setSwagger(r *gin.RouterGroup) {
	if settings.C().App.Swagger.Enable {
		swagger.SwaggerInfo.Version = settings.C().App.Version
		swagger.SwaggerInfo.Title = settings.C().App.Name
		swagger.SwaggerInfo.BasePath = r.BasePath()
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
