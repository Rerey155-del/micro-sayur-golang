package config

import (
	"log"

	"github.com/spf13/viper"
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
	App  App
	DB   DBConfig
}

func NewConfig() *Config {
	// Set default values or read from .env
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),
		},
		DB: DBConfig{
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetString("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
	}
}
