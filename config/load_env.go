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
	}
	return nil
}
