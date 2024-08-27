package jwt_utils

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (*entity.TokenClaims, error) {
	secretKey := []byte(config.Env.JWT_Secret)

	claims := &entity.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
