package config

import (
	"log"
	"os"
	"strconv"
)

type App struct {
	AppPort string
	AppEnv  string
}

type DBConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	DBMaxOpen int
	DBMaxIdle int
}

type Config struct {
	App App
	DB  DBConfig
}

func NewConfig() *Config {
	if _, err := os.Stat(".env"); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	return &Config{
		App: App{
			AppPort: getEnv("APP_PORT", "8082"),
			AppEnv:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:      getEnv("DATABASE_HOST", "127.0.0.1"),
			Port:      getEnv("DATABASE_PORT", "3306"),
			User:      getEnv("DATABASE_USER", "root"),
			Password:  getEnv("DATABASE_PASSWORD", ""),
			DBName:    getEnv("DATABASE_NAME", "sayur-order-service"),
			DBMaxOpen: getEnvAsInt("DATABASE_MAX_OPEN_CONNECTION", 100),
			DBMaxIdle: getEnvAsInt("DATABASE_MAX_IDLE_CONNECTION", 10),
		},
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
