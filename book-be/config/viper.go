package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	workingDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Warning: Failed to load .env file, continuing with environment variables.")
		}
	} else {
		fmt.Println("No .env file found, skipping loading of .env.")
	}

	config.AddConfigPath(workingDir)
	config.Set("app.name", os.Getenv("APP_NAME"))
	config.Set("app.version", os.Getenv("APP_VERSION"))
	config.Set("app.port", os.Getenv("APP_PORT"))
	config.Set("db.host", os.Getenv("DB_HOST"))
	config.Set("db.name", os.Getenv("DB_NAME"))
	config.Set("db.port", os.Getenv("DB_PORT"))
	config.Set("db.username", os.Getenv("DB_USERNAME"))
	config.Set("db.password", os.Getenv("DB_PASSWORD"))
	config.Set("jwt.secret", os.Getenv("JWT_SECRET"))

	return config
}
