package middleware

import (
	"gin_template/global"

	"github.com/gin-gonic/gin"
)

func InjectDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 注入数据库实例到上下文
		c.Set("db", global.DB)
		c.Next()
	}
}
