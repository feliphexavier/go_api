package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJWT      string

	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env")
	}
	log.Println("Config Loaded")
	return &Config{
		Port:           os.Getenv("PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		SecretJWT:      os.Getenv("SECRET_JWT"),
		DBUser:         os.Getenv("POSTGRES_USER"),
		DBPassword:     os.Getenv("POSTGRES_PASSWORD"),
		DBName:         os.Getenv("POSTGRES_DB"),
		DBHost:         os.Getenv("POSTGRES_HOST"),
		DBPort:         os.Getenv("POSTGRES_PORT"),
	}, nil
}
