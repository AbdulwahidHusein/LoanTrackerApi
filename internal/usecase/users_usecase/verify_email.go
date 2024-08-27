package users_usecase

import (
	"LoanTrackerApi/config"
	"LoanTrackerApi/internal/entity"
	"LoanTrackerApi/pkg/email_utils"
	"LoanTrackerApi/pkg/jwt_utils"
	"context"
	"errors"
	"fmt"
	"log"
)

func (u *Usecase) RequestEmailVerification(user entity.LoginUserDTO) error {
	var emailSender email_utils.EmailSender

	dbUSer, err := u.userRepo.FindByEmail(context.Background(), user.Email)

	if err != nil {
		return err
	}

	if dbUSer == nil {
		return errors.New("user with this email does not exist")
	}
	Config := config.Env

	emailProvider := Config.EMAIL_PROVIDER

	fmt.Println("Port: ", Config.EMAIL_PORT)
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

	go func() {
		claims := entity.TokenClaims{
			Email:  user.Email,
			Role:   "email-verification",
			UserId: "email-verification",
		}
		token, err := jwt_utils.CreateToken(claims)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
		}
		err = emailSender.SendVerificationEmail(user.Email, token)
		if err != nil {
			log.Printf("Failed to send verification email: %v", err)
		}
	}()

	return nil
}

func (u *Usecase) VerifyEmail(token string) error {
	claims, err := jwt_utils.ValidateToken(token)
	if err != nil {
		return err
	}

	if claims.Role != "email-verification" || claims.UserId != "email-verification" {
		return errors.New("invalid token")
	}

	issuerEmail := claims.Email
	user, err := u.userRepo.FindByEmail(context.Background(), issuerEmail)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user with this email does not exist")
	}
	// user. = true
	err = u.userRepo.Verify(context.Background(), user.ID.Hex())

	if err != nil {
		return err
	}
	return nil

}
