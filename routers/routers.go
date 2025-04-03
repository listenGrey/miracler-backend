package routers

import (
	"alpha/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}        // 允许前端的地址
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"} // 允许的方法
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	// 使用zap
	r.Use(
		cors.New(config),
		//logger.GinLogger(),
		//logger.GinRecovery(false),
		//Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		//middlewares.RateLimitMiddleware(2*time.Second, 40), // 每两秒钟添加十个令牌  全局限流
	)

	v1 := r.Group("/api/v1")
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)
	// v1.GET("/refresh_token", controller.RefreshTokenHandler)

	// 中间件
	v1.Use(middlewares.JWTAuthMiddleWare()) //JWT认证
	{
		// Today
		v1.GET("/today", controllers.Today)
		v1.POST("/summary", controllers.Summary)

		// Calendar
		v1.GET("/calendar", controllers.Calendar)
		v1.PUT("/one-day", controllers.ModifyOneDay)

		// New
		v1.POST("/event", controllers.NewEvent)

		// Overview
		v1.GET("/overview", controllers.Overview)

		// Profile 修改事件和项目
		v1.GET("/profile", controllers.Profile)
		v1.PUT("/event", controllers.ModifyEvent)
		v1.DELETE("/event", controllers.DeleteEvent)

		// Testing
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}
	r.NoRoute(func(c *gin.Context) {
		errHandler.ResponseError(c, code.NotFound)
	})
	return r
}
