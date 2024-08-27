package entity

import "github.com/golang-jwt/jwt"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
