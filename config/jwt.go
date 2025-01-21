package config

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTSecretKey = []byte("SecretKey")

// GenerateToken 生成token
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecretKey)
}

// ParseToken 解析token
func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["userID"].(float64)), nil
	}
	return 0, err
}
