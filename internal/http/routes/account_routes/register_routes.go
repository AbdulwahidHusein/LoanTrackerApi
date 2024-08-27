package account_routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	handler := GetHandler()
	usersGroup := router.Group("/users")
	usersGroup.POST("/register", handler.RegisterationHandler)
	usersGroup.POST("/login", handler.LoginHandler)
	usersGroup.POST("/verify-email/request", handler.RequestVerifyEmail)
	usersGroup.POST("/verify-email", handler.VerifyEmail)
	return router
}
