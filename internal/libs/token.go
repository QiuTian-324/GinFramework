package libs

import (
	"gin_template/pkg"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	ID                   uint   `json:"id"`
	Username             string `json:"nickname"`
	Role                 int    `json:"role"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(id uint, username string) (string, error) {
	// viper获取jwt密钥
	var customSecret = []byte(viper.GetString("jwt.secret"))

	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(viper.GetInt64("jwt.expires_in")))),
			// 签发人
			Issuer: viper.GetString("jwt.issuer"),
		},
	}
	// 生产jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token字符串
	return token.SignedString(customSecret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	// viper获取jwt密钥
	var customSecret = []byte(viper.GetString("jwt.secret"))
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return customSecret, nil
	})

	if err != nil {
		pkg.Error("解析token失败", err)
		return nil, err
	}

	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
