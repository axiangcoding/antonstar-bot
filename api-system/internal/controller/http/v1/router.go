package v1

import (
	swagger "github.com/axiangcoding/antonstar-bot/api"
	http2 "github.com/axiangcoding/antonstar-bot/internal/controller/http"
	"github.com/axiangcoding/antonstar-bot/internal/controller/middleware"
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

// InitRouter
// @title        axiangcoding/anton-star
// @version      1.0.0
// @description  api system build by ax-web
// @termsOfService
// @contact.name  axiangcoding
// @contact.url
// @contact.email  axiangcoding@gmail.com
// @license.name
// @license.url
// @accept                      json
// @produce                     json
// @securityDefinitions.apikey  AppToken
// @in                          query
// @name                        app_token
// @securityDefinitions.apikey  CqhttpSelfID
// @in                          header
// @name                        X-Self-ID
// @securityDefinitions.apikey  CqhttpSignature
// @in                          header
// @name                        X-Signature
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
	base := r.Group(setting.C().Server.BasePath)
	setWebResources(base)
	setRouterApiV1(base)
	return r
}

func setWebResources(r *gin.RouterGroup) {
	r.GET("/", http2.RootRedirect)
	r.StaticFS("/web", http.Dir("web"))
}

func setRouterApiV1(r *gin.RouterGroup) {
	api := r.Group("/api")
	setSwagger(api)
	groupV1 := api.Group("/v1")
	{
		system := groupV1.Group("/system")
		{
			system.GET("/info", SystemInfo)
		}
		app := groupV1.Group("/app")
		{
			app.GET("/info", AppInfo)
		}
		cqhttp := groupV1.Group("/cqhttp")
		{
			cqhttpAuth := middleware.CqhttpAuth(
				setting.C().App.Service.CqHttp.SelfQQ,
				setting.C().App.Service.CqHttp.Secret)
			cqhttp.POST("/receive/event", cqhttpAuth, CqHttpReceiveEvent)
			cqhttp.GET("/status", CqHttpStatus)
		}
		wt := groupV1.Group("/wt")
		{
			wt.GET("/profile", GameUserProfile)
			wt.POST("/profile/update", UpdateGameUserProfile)
		}
		mission := groupV1.Group("/mission")
		{
			mission.GET("/", GetMission)
		}
	}
}

func setSwagger(r *gin.RouterGroup) {
	if setting.C().App.Swagger.Enable {
		swagger.SwaggerInfo.Version = setting.C().App.Version
		swagger.SwaggerInfo.Title = setting.C().App.Name
		swagger.SwaggerInfo.BasePath = r.BasePath()
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
