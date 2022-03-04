package router

import (
	"axiangcoding/antonstar/api-system/controller/middleware"
	"axiangcoding/antonstar/api-system/controller/v1"
	"axiangcoding/antonstar/api-system/settings"
	"axiangcoding/antonstar/api-system/swagger"
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

// 设置cors头
func setCors(r *gin.Engine) {
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AddAllowMethods("OPTIONS")
	// config.AddAllowHeaders(app.AuthHeader)
	// r.Use(cors.New(config))
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
		user := groupV1.Group("/user")
		{
			user.POST("/login", middleware.CaptchaCheck(), v1.UserLogin)
			user.POST("/register", middleware.CaptchaCheck(), v1.UserRegister)
			user.POST("/logout", middleware.AuthCheck(), v1.UserLogout)
			user.POST("/value/exist", v1.IsKeyFieldValueExist)
			user.POST("/info", v1.UserInfo)
		}
		system := groupV1.Group("/system")
		{
			system.GET("/info", v1.SystemInfo)
		}
		site := groupV1.Group("/site")
		{
			site.GET("/notice/last", v1.GetLastSiteNotice)
			site.POST("/notice/", middleware.AuthCheck(), v1.PostSiteNotice)
		}
		visit := groupV1.Group("/visits")
		{
			visit.POST("/visit", v1.PostVisit)
			visit.GET("/", middleware.AuthCheck(), v1.GetVisits)
			visit.GET("/count", v1.GetVisitCount)
		}
		captcha := groupV1.Group("/captcha")
		{
			captcha.GET("/", v1.GenerateCaptcha)
			captcha.GET("/:file", v1.GetCaptcha)
			captcha.POST("/verify")
		}
		warThunder := groupV1.Group("/war_thunder")
		{
			warThunder.GET("/userinfo/queries", v1.GetUserInfoQueries)
			warThunder.POST("/userinfo/refresh", v1.PostUserInfoRefresh)
			warThunder.GET("/userinfo", v1.GetUserInfo)
			warThunder.GET("/userinfo/query/count", v1.GetQueryCount)
			// warThunder.StaticFile("/mock.html", "./resources/index.html")
		}
	}
}
