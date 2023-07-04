package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("A87FF679A2F3E71D9181A67B7542122CZ")

// CreateToken 创建一个JWT格式的Token
func CreateToken(data map[string]string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(duration).Unix(),
		"custom": data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// VerifyToken 验证一个JWT格式的Token
func VerifyToken(tokenString string) (bool, map[string]interface{}) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signingKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return ok, claims["custom"].(map[string]interface{})
	}
	return false, nil
}
