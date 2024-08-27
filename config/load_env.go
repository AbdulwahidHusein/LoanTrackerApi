package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DatabaseUrl          string
	JWT_Secret           string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	EMAIL_PROVIDER       string
	SMTP_HOST            string
	EMAIL_PORT           string
	SENDER_EMAIL         string
	SENDER_PASSWORD      string
}

var Env EnvConfig

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	Env = EnvConfig{
		DatabaseUrl:          os.Getenv("DATABASE_URL"),
		JWT_Secret:           os.Getenv("JWT_SECRET"),
		GOOGLE_CLIENT_ID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET: os.Getenv("GOOGLE_CLIENT_SECRET"),
		EMAIL_PROVIDER:       os.Getenv("EMAIL_PROVIDER"),
		SMTP_HOST:            os.Getenv("SMTP_HOST"),
		EMAIL_PORT:           os.Getenv("EMAIL_PORT"),
		SENDER_EMAIL:         os.Getenv("SENDER_EMAIL"),
		SENDER_PASSWORD:      os.Getenv("SENDER_PASSWORD"),
	}
	return nil
}
