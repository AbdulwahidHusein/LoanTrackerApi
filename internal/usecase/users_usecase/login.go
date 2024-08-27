package users_usecase

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/pkg/jwt_utils"
	"LoanTrackerApi/pkg/validators"
	"context"
	"errors"
)

func (u *Usecase) Login(ctx context.Context, user *entity.LoginUserDTO) (entity.Token, error) {
	email := user.Email
	password := user.Password

	dbUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return entity.Token{}, errors.New("some error occurred")
	}
	if dbUser == nil {
		return entity.Token{}, errors.New("user not found")
	}

	if !validators.CheckPasswordHash(password, dbUser.Password) {
		config.Logger.AddLog(ctx, "Login", "user with email "+email+" failed provided password does not match")
		return entity.Token{}, errors.New("invalid password")
	}
	claims := entity.TokenClaims{
		Email:  dbUser.Email,
		UserId: dbUser.ID.Hex(),
		Role:   dbUser.Role,
	}

	access, err := jwt_utils.CreateToken(claims)
	if err != nil {
		config.Logger.AddLog(ctx, "Login", "some error occurred while creating access token for user with email "+email)
		return entity.Token{}, errors.New("some error occurred")
	}
	refresh, err := jwt_utils.CreateToken(claims)
	if err != nil {
		config.Logger.AddLog(ctx, "Login", "some error occurred while creating refresh token for user with email "+email)
		return entity.Token{}, errors.New("some error occurred")
	}

	config.Logger.AddLog(ctx, "Login", "user with email "+email+" logged in")

	return entity.Token{AccessToken: access, RefreshToken: refresh}, nil
}
