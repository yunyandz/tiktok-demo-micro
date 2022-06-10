package jwtx

import "github.com/golang-jwt/jwt"

type UserInfo struct {
	Username string `json:"username"`
	UserId   uint64 `json:"user_id"`
}

type UserClaims struct {
	UserInfo
	jwt.StandardClaims
}
