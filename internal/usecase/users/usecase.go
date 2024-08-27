package users

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, user *entity.LoginUserDTO) (entity.Token, error)
	// ForgotPassword(ctx context.Context, email string) error
	// VerifyEmail(ctx context.Context, token string) error
	// UpdateProfile(ctx context.Context, user *entity.User) error
	// GoogleLogin(ctx context.Context, user *entity.User) error
}
