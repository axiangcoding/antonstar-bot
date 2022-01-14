package router

import (
	"axiangcoding/antonstar/api-system/api/docs"
	v1 "axiangcoding/antonstar/api-system/api/v1"
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 全局中间件
	// 使用自定义中间件
	r.Use(middleware.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	setCors(r)
	setRouterV1(r)
	return r
}

func setCors(r *gin.Engine) {
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AddAllowMethods("OPTIONS")
	// config.AddAllowHeaders(app.AuthHeader)
	// r.Use(cors.New(config))
}

func setSwagger(r *gin.RouterGroup) {
	if conf.Config.App.Swagger.Enable {
		docs.SwaggerInfo.Version = conf.Config.App.Version
		docs.SwaggerInfo.Title = conf.Config.App.Name
		docs.SwaggerInfo.BasePath = conf.Config.Server.BasePath
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setRouterV1(r *gin.Engine) {
	base := r.Group(conf.Config.Server.BasePath)
	setSwagger(base)
	groupV1 := base.Group("/v1")
	{
		demo := groupV1.Group("/demo")
		{
			demo.GET("/get", v1.DemoGet)
			demo.POST("/post", v1.DemoPost)
		}
		test := groupV1.Group("/test")
		{
			test.GET("/test-log", v1.TestLog)
		}
		user := groupV1.Group("/user")
		{
			user.POST("/login", v1.UserLogin)
			user.POST("/register", v1.UserRegister)
			user.POST("/logout", middleware.AuthCheck(), v1.UserLogout)
		}
		system := groupV1.Group("/system", middleware.AuthCheck())
		{
			system.GET("/info", v1.SystemInfo)
		}
		visit := groupV1.Group("/visits")
		{
			visit.POST("/visit", v1.PostVisit)
			visit.GET("/", middleware.AuthCheck(), v1.GetVisits)
			visit.GET("/count", v1.GetVisitCount)
		}
		warThunder := groupV1.Group("/war_thunder")
		{
			warThunder.GET("/userinfo/queries", v1.GetUserInfoQueries)
			warThunder.POST("/userinfo/refresh", v1.PostUserInfoRefresh)
			warThunder.GET("/userinfo", v1.GetUserInfo)
			// warThunder.StaticFile("/mock.html", "./resources/index.html")
		}
	}
}
