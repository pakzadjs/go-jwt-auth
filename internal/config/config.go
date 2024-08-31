package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	Port         string
	JWTSecret    string
	AllowOrigins []string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	}
}
