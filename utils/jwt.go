package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mojocn/turbo-iot/config"
	"time"
)

func JwtGenerateToken(jwtAuthRedisKey string, expiresAfterTime time.Duration) (string, error) {
	stdClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiresAfterTime).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        jwtAuthRedisKey,
		Issuer:    config.AppName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.JwtSignSecret))
	return tokenString, err
}

func JwtParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JwtSignSecret), nil
	})
	return &claims, err
}
