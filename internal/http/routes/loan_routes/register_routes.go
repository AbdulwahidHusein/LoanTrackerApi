package loan_routes

import (
	"LoanTrackerApi/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterLoanRoutes(route *gin.Engine) {
	handler := GetHandler()

	authGroup := route.Group("/loan", middlewares.AuthMiddleware())
	authGroup.POST("/apply", handler.Apply)
	authGroup.GET("/:id", handler.GetLoanHandler)

	adminGroup := route.Group("/admin/loan", middlewares.AuthMiddleware())

	adminGroup.GET("/", handler.GetLoans)
	adminGroup.PATCH("approve/:id", handler.ApproveLoanHanlder)
	adminGroup.PATCH("reject/:id", handler.RejectLoanHandler)
	adminGroup.DELETE("/:id", handler.DeleteLoanHandler)

}
