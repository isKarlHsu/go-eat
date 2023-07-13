package jwt

import (
	"eat/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type JwtPayLoad struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func GenToken(user JwtPayLoad) (string, error) {
	var secretKey = []byte(global.Config.Jwt.Secret)
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	var secretKey = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("token parse error: %s", err))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
