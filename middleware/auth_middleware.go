package middleware

import (
	"context"
	"fmt"
	"gin_template/global"
	"gin_template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthToken token校验
func AuthToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	ip := c.ClientIP()
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.BadWithMessage(fmt.Sprintf("非法请求，请停止任何非法操作，ip%s已记录！", ip)))
		return
	}
	//  token 存在
	//  尝试解密token
	currentUser, err := utils.ValidationToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.BadWithMessage(fmt.Sprintf("非法请求，请停止任何非法操作，ip%s已记录！", ip)))
		return
	}
	// redis 获取token
	result, _ := global.RedisClient.Get(context.Background(), global.RedisKey+currentUser).Result()
	if result == "" {
		// 不存在用户信息
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.BadWithMessage(fmt.Sprintf("非法请求，请停止任何非法操作，ip%s已记录！", ip)))
		return
	}
	// 用户名设置到上下文中
	c.Set("user", currentUser)
	c.Next()
}
