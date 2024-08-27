package account_routes

import (
	"LoanTrackerApi/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	handler := GetHandler()
	usersGroup := router.Group("/users")
	usersGroup.POST("/register", handler.RegisterationHandler)
	usersGroup.POST("/login", handler.LoginHandler)
	usersGroup.POST("/verify-email/request", handler.RequestVerifyEmail)
	usersGroup.GET("/verify-email", handler.VerifyEmail)
	usersGroup.POST("/reset-password/request", handler.ResetPasswordRequest)
	usersGroup.POST("/reset-password", handler.ResetPassword)
	usersGroup.POST("/token/refresh", handler.Refresh)
	usersGroup.GET("/profile", middlewares.AuthMiddleware(), handler.GetMyProfile)

	adminGroup := router.Group("/admin", middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	adminGroup.GET("/users", handler.AdminGetUsers)
	adminGroup.DELETE("/users/:id", handler.AdminDeleteUser)
	// adminGroup.GET("/users/:id", middlewares.AuthMiddleware(), handler.GetUser)

}
