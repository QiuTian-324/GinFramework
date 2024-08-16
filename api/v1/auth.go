package v1

import (
	"context"
	"fmt"
	"gin_template/global"
	"gin_template/model/user"
	"gin_template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 用户登录结构体
type login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// LoginApi 登录接口
func LoginApi(c *gin.Context) {
	loginUser := login{}
	err := c.BindJSON(&loginUser)
	if err != nil {
		// 参数出问题了
		c.AbortWithStatusJSON(http.StatusOK, utils.FailWithMessage("参数错误"))
		return
	}
	userTemp := user.AuthUser{}
	err = global.DB.Select([]string{"username", "is_use"}).Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&userTemp).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.FailWithMessage("用户名或密码错误，请重试"))
		return
	} else {
		// 用户名和密码正确
		if userTemp.IsUse == false {
			c.AbortWithStatusJSON(http.StatusOK, utils.FailWithMessage("用户已被禁用，请联系站点管理员"))
			return
		} else {
			// 用户没有被锁定
			global.Logger.Infof(fmt.Sprintf("用户:%s  登录成功!", userTemp.Username))
			// 生成token
			token, err := utils.CreateToken(userTemp.Username)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, utils.FailWithMessage("服务器内部出现问题，请联系站点管理员"))
				return
			}
			// 存入redis
			err = global.RedisClient.SetEx(context.Background(), global.RedisKey+userTemp.Username, token, time.Hour*2).Err()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, utils.FailWithMessage("服务器内部出现问题，请联系站点管理员"))
				return
			}
			makes := make(map[string]interface{})
			makes["Authorization"] = token
			c.JSON(http.StatusOK, utils.OKWithData(makes))
			return
		}

	}
}
