/*
 * @Author: charley zhu
 * @Date: 2023-10-11 10:02:53
 * @LastEditTime: 2023-10-13 07:12:05
 * @LastEditors: charley zhu
 * @Description:
 */
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type customedClaims struct {
	UserName string `json:"username"`
	UserId   string `json:"user_id"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("custom_secret")

/**
 * @description: gen token
 * @param {*} username
 * @param {string} user_id
 * @return {*}
 */
func GenToken(username, user_id string) (string, error) {
	claims := customedClaims{
		username,
		user_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "global_backend",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

/**
 * @description: parse token
 * @param {string} tokenString
 * @return {*}
 */
func ParseToken(tokenString string) (*customedClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customedClaims{}, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*customedClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
