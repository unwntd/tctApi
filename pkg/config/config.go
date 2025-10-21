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
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
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
	}
}
