package account_routes

import (
	"LoanTrackerApi/internal/http/handlers/account_handler"
	"LoanTrackerApi/internal/repository/mongodb/user_repo"
	"LoanTrackerApi/internal/usecase/users_usecase"
	"LoanTrackerApi/pkg/mongodb"
)

func GetHandler() *account_handler.Handler {
	userCollection := mongodb.GetCollection("users")
	userRepo := user_repo.NewMongoUserRepository(userCollection)
	userUsecase := users_usecase.NewUsecase(userRepo)
	return account_handler.NewHandler(userUsecase)
}
