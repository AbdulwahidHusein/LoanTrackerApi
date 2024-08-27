package account_routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	handler := GetHandler()
	usersGroup := router.Group("/users")
	usersGroup.POST("/register", handler.RegisterationHandler)
	usersGroup.POST("/login", handler.LoginHandler)
	return router
}
