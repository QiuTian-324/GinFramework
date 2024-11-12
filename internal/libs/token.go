package libs

import (
	"gin_template/pkg"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// CustomClaims 自定义声明类型
type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

// GenAccessToken 生成 Access Token
func GenAccessToken(id uint, username string) (string, error) {
	secret := []byte(viper.GetString("jwt.secret"))
	claims := CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(viper.GetInt64("jwt.access_token_expires_in")))),
			Issuer:    viper.GetString("jwt.issuer"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// GenRefreshToken 生成 Refresh Token
func GenRefreshToken(id uint, username string) (string, error) {
	secret := []byte(viper.GetString("jwt.secret"))
	claims := CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(viper.GetInt64("jwt.refresh_token_expires_in")))),
			Issuer:    viper.GetString("jwt.issuer"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 解析 Token (支持 Access 和 Refresh Token)
func ParseToken(tokenStr string) (*CustomClaims, error) {
	secret := []byte(viper.GetString("jwt.secret"))
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		pkg.Error("解析 token 失败", err)
		return nil, err
	}

	if tokenClaims != nil {
		if c, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
