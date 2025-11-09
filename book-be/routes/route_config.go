package routes

import (
	"book-be/handler"
	"book-be/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App   *echo.Echo
	DB    *gorm.DB
	Viper *viper.Viper
	Log   *logrus.Logger

	AuthHandler      *handler.AuthHandler
	BookHandler      *handler.BookHandler
	AuthorHandler    *handler.AuthorHandler
	PublisherHandler *handler.PublisherHandler

	AuthMiddleware *middleware.AuthMiddleware
}

func (c *RouteConfig) Setup() {
	c.GuestRoute()
	c.AuthRoute()
	c.BookRoute()
	c.PublisherRoute()
	c.AuthorRoute()
}
