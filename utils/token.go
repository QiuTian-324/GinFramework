package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func CreateToken(user string) (string, error) {
	mySigningKey := []byte("Key1234567890")

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"buding",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "admin",
			Subject:   user,
			ID:        "-100",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	return ss, err

}
func ValidationToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Key1234567890"), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims.RegisteredClaims.Subject, nil
	} else {
		return "", errors.New("cant parse claims")
	}
}
