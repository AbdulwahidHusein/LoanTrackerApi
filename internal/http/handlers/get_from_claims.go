package handlers

import (
	"LoanTrackerApi/internal/entity"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*entity.TokenClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("user not found")
	}
	userClaims, ok := claims.(*entity.TokenClaims)
	if !ok {
		return nil, errors.New("user not found")
	}
	return userClaims, nil
}

func GetUserId(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}

func GetEmail(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}

func GetRole(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
