package users_usecase

import (
	"LoanTrackerApi/internal/entity"
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, user *entity.LoginUserDTO) (entity.Token, error)
	RequestEmailVerification(user entity.LoginUserDTO) error
	VerifyEmail(token string) error
	RequestPasswordResetUsecase(userEmail string) error
	ResetPassword(resetToken string, newPassword string) error
	GetMyProfile(ctx context.Context, id string) (entity.GetUserDTO, error)
	AdminGetAllUsers(ctx context.Context, page int, pageSize int) ([]*entity.GetUserDTO, error)
	AdminDeleteUser(ctx context.Context, id string) error
	// ForgotPassword(ctx context.Context, email string) error
	// VerifyEmail(ctx context.Context, token string) error
	// UpdateProfile(ctx context.Context, user *entity.User) error
	// GoogleLogin(ctx context.Context, user *entity.User) error
}
