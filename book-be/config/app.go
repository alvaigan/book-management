package config

import (
	"book-be/handler"
	"book-be/middleware"
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

	// handlers
	authHandler := handler.NewAuthHandler(ac.App, ac.DB, ac.Viper, ac.Log)
	bookHandler := handler.NewBookHandler(ac.App, ac.DB, ac.Viper, ac.Log)
	authorHandler := handler.NewAuthorHandler(ac.App, ac.DB, ac.Viper, ac.Log)
	publisherHandler := handler.NewPublisherHandler(ac.App, ac.DB, ac.Viper, ac.Log)

	// middlewares
	authMiddleware := middleware.NewAuthMiddleware(ac.DB)

	router := routes.RouteConfig{
		App:   ac.App,
		DB:    ac.DB,
		Viper: ac.Viper,
		Log:   ac.Log,

		AuthHandler:      authHandler,
		BookHandler:      bookHandler,
		AuthorHandler:    authorHandler,
		PublisherHandler: publisherHandler,

		AuthMiddleware: authMiddleware,
	}

	router.Setup()
}
