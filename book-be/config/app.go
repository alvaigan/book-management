package config

import (
	"book-be/handler"
	"book-be/routes"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	App   *echo.Echo
	DB    *gorm.DB
	Viper *viper.Viper
	Log   *logrus.Logger
}

func NewApp(ac *AppConfig) {

	handler := handler.NewHandler(ac.App, ac.DB, ac.Viper, ac.Log)

	router := routes.RouteConfig{
		App:     ac.App,
		DB:      ac.DB,
		Viper:   ac.Viper,
		Log:     ac.Log,
		Handler: handler,
	}

	router.Setup()
}
