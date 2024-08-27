package main

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/http/routes/account_routes"
	"LoanTrackerApi/pkg/mongodb"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
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

	router.Run()
}
