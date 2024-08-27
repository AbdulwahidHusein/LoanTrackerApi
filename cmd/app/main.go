package main

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/http/routes/account_routes"
	"LoanTrackerApi/internal/http/routes/loan_routes"
	"LoanTrackerApi/pkg/mongodb"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetupLogger()
	err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = mongodb.ConnectDB(config.Env.DatabaseUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mongodb.DisconnectDB()
	mongodb.InitializeCollections()

	router := gin.Default()

	account_routes.RegisterRoutes(router)
	loan_routes.RegisterLoanRoutes(router)

	router.Run()
}
