package middleware

import (
	"fmt"
	"strings"
	"yug_server/global"
	"yug_server/internal/libs"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
)

// AuthToken token校验
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先从请求头获取 token
		token := c.Request.Header.Get("Authorization")

		// 如果请求头中没有 token，尝试从查询参数中获取
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			ip := c.ClientIP()
			pkg.Error("token 不存在", nil)
			libs.BadRequestResponse(c, fmt.Sprintf("非法请求，请停止任何非法操作，IP %s 已记录！", ip))
			return
		}

		// 分割 token
		parts := strings.SplitN(token, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			token = parts[1]
		}

		ip := c.ClientIP()
		// 尝试解密 token
		currentUser, err := libs.ParseToken(token)
		if err != nil {
			pkg.Error("token 解密失败", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，解密 token 失败，IP %s 已记录！", ip))
			return
		}

		// 检查 Redis 中是否存在 token
		redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, currentUser.ID)
		_, err = libs.RedisGet(c, redisKey)
		if err != nil {
			pkg.Error("Redis 中不存在 token", err)
			libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，用户未认证，IP %s 已记录！", ip))
			return
		}

		// 将用户名设置到上下文中
		c.Set("id", currentUser.ID)
		c.Next()
	}
}
