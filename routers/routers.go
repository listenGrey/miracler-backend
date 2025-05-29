package routers

import (
	"alpha/handler"
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
	config.AllowOrigins = []string{"http://localhost:3000"}                 // 允许前端的地址
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT", "PATCH"} // 允许的方法
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	// 使用zap
	r.Use(
		cors.New(config),
	)

	v1 := r.Group("/api/v1")
	v1.POST("/register", handler.Register)
	v1.POST("/login", handler.Login)
	// v1.GET("/refresh_token", controller.RefreshTokenHandler)

	// 中间件
	// v1.Use(middlewares.JWTAuthMiddleWare()) //JWT认证
	{
		// Today
		v1.GET("/today", handler.Today)
		v1.POST("/summary", handler.Summary)
		v1.POST("/layup", handler.Layup)

		// Event
		v1.GET("/events", handler.GetEvents)
		v1.PATCH("/events/:eventId/items/:itemIndex", handler.UpdateItem)

		// Calendar
		v1.GET("/calendar", handler.Calendar)
		v1.PUT("/one-day", handler.ModifyOneDay)

		// New
		v1.POST("/events", handler.NewEvent)

		// Overview
		v1.GET("/overview", handler.Overview)

		// Milestone 打算把日历放在总览的里面，里程碑记录完成情况和线条路程
		v1.GET("/milestone")

		// Profile 修改事件和项目
		v1.GET("/profile", handler.Profile)
		v1.PUT("/event", handler.ModifyEvent)
		v1.DELETE("/events", handler.DeleteEvent)

		// Testing
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}
	/*r.NoRoute(func(c *gin.Context) {
		errHandler.ResponseError(c, code.NotFound)
	})*/
	return r
}
