package jwt_utils

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(claims entity.TokenClaims) (string, error) {
	secret := []byte(config.Env.JWT_Secret)

	ACCESS_TOKEN_LIFE_TIME := 2 * 24 * time.Hour

	expirationTime := time.Now().Add(ACCESS_TOKEN_LIFE_TIME)
	if claims.Role == "" {
		claims.Role = "user" // Set default role to "user" if no role is passed
	}

	claims.ExpiresAt = expirationTime.Unix()

	secretKey := secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err // Return an empty string and the error
	}

	return signedToken, nil // Return the signed token and no error

}
