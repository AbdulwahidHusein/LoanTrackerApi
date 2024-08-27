package users_usecase

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/repository"
	"LoanTrackerApi/pkg/validators"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usecase struct {
	userRepo repository.UserRepository
}

func NewUsecase(userRepo repository.UserRepository) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}

func (u *Usecase) Register(ctx context.Context, user *entity.User) error {
	email := user.Email
	password := user.Password

	u.RequestEmailVerification(entity.LoginUserDTO{
		Email: email})

	dbUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if dbUser != nil {
		return errors.New("user already exists")
	}
	user.Password, err = validators.HashPassword(password)
	if err != nil {
		return err
	}
	user.ID = primitive.NewObjectID()

	return u.userRepo.Create(ctx, user)
}
