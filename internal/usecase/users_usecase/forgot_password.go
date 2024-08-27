package users_usecase

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/pkg/email_utils"
	"LoanTrackerApi/pkg/jwt_utils"
	"LoanTrackerApi/pkg/validators"
	"context"
	"errors"
	"fmt"
	"log"
)

func (u *Usecase) RequestPasswordResetUsecase(userEmail string) error {
	var emailSender email_utils.EmailSender
	Config := config.Env

	emailProvider := Config.EMAIL_PROVIDER

	switch emailProvider {
	case "simple":
		emailSender = email_utils.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.SENDER_EMAIL,
			Config.SENDER_PASSWORD,
		)
	default:
		emailSender = email_utils.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.SENDER_EMAIL,
			Config.SENDER_PASSWORD,
		)
	}
	claims := entity.TokenClaims{
		Email:  userEmail,
		Role:   "password-reset",
		UserId: "password-reset",
	}
	token, err := jwt_utils.CreateToken(claims)
	if err != nil {
		return err
	}
	accessToken := token
	go func() {
		err := emailSender.SendPasswordResetEmail(userEmail, accessToken)
		if err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	return nil
}

func (u *Usecase) ResetPassword(resetToken string, newPassword string) error {
	claims, err := jwt_utils.ValidateToken(resetToken)

	if err != nil {
		return err
	}
	fmt.Println(claims, "claims ====================")
	if claims.Role != "password-reset" || claims.UserId != "password-reset" {
		return errors.New("invalid token")
	}
	email := claims.Email
	hashedPassword, err := validators.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user, err := u.userRepo.FindByEmail(context.Background(), email)

	if user != nil && err == nil {
		user.Password = hashedPassword
		err = u.userRepo.ChangePassword(context.Background(), user, hashedPassword)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
