package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName   string
	AppPort   string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	JWTSecret string
	AppMode   string
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	// Set default values
	viper.SetDefault("APP_NAME", "tctAPI")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_NAME", "tctapi")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("JWT_SECRET", "your-secret-key")
	viper.SetDefault("APP_MODE", "debug")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %s. Using default values.", err)
	}

	return &Config{
		AppName:   viper.GetString("APP_NAME"),
		AppPort:   viper.GetString("APP_PORT"),
		DBHost:    viper.GetString("DB_HOST"),
		DBUser:    viper.GetString("DB_USER"),
		DBPass:    viper.GetString("DB_PASS"),
		DBName:    viper.GetString("DB_NAME"),
		DBPort:    viper.GetString("DB_PORT"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		AppMode:   viper.GetString("APP_MODE"),
	}
}
