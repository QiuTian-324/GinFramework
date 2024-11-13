package middleware

import (
	"context"
	"fmt"
	"gin_template/global"
	"gin_template/internal/libs"

	"github.com/gin-gonic/gin"
)

// AuthToken token校验
func AuthToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	ip := c.ClientIP()

	// 检查 token 是否存在
	if token == "" {
		libs.BadRequestResponse(c, fmt.Sprintf("非法请求，请停止任何非法操作，IP %s 已记录！", ip))
		return
	}

	// 尝试解密 token
	currentUser, err := libs.ParseToken(token)
	if err != nil {
		libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，解密 token 失败，IP %s 已记录！", ip))
		return
	}

	// 检查 Redis 中是否存在 token
	result, _ := global.RedisClient.Get(context.Background(), global.RedisPrefix+currentUser.Username).Result()
	if result == "" {
		libs.UnauthorizedResponse(c, fmt.Sprintf("非法请求，用户未认证，IP %s 已记录！", ip))
		return
	}

	// 将用户名设置到上下文中
	c.Set("user", currentUser)
	c.Next()
}
