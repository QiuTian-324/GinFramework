package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		// 设置 CORS 头
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Sec-Ch-Ua-Mobile,Sec-Ch-Ua,Sec-Ch-Ua-Platform,Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 设置 Content-Security-Policy 头
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src 'self' http://127.0.0.1:9000;")

		// 放行所有 OPTIONS 方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
