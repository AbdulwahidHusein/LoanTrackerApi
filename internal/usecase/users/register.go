package users

import (
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/internal/repository"
	"LoanTrackerApi/pkg/validators"
	"context"
	"errors"
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

	return u.userRepo.Create(ctx, user)
}
