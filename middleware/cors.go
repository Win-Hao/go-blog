package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	// 创建CORS中间件的配置
	config := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length", "Authorization"},
		//AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//    return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}

	// 使用配置创建CORS中间件
	corsMiddleware := cors.New(config)

	// 返回一个函数，该函数将中间件应用到每个请求
	return func(c *gin.Context) {
		// 调用中间件处理函数
		corsMiddleware(c)

		// 如果请求是OPTIONS方法，处理预检请求
		if c.Request.Method == "OPTIONS" {
			// 并且没有设置其他中间件
			c.AbortWithStatus(204)
		} else {
			// 继续处理请求
			c.Next()
		}
	}
}
