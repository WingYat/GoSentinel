/**
 * @Auth: yycheung - win9yat.cheung@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 9/11/2023 - 11:40 pm
 * @File: internal/router/router.go
 * @Desc:
 */

package router

import (
	"GoSentinel/api/controller"
	"GoSentinel/api/middleware"
	"github.com/gin-gonic/gin"
	// 其他导入...
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// 全局中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// JWT中间件
	authMiddleware := middleware.JwtMiddleware()

	// API路由组
	apiGroup := router.Group("/api")
	apiGroup.Use(authMiddleware)
	{
		apiGroup.GET("/service", controller.GetService)
		apiGroup.POST("/service", controller.CreateService)
		// 更多API路由...
	}

	// 监控路由组
	monitorGroup := router.Group("/monitor")
	{
		monitorGroup.GET("/metrics", controller.GetMetrics)
		// 更多监控路由...
	}

	return router
}
