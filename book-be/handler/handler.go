package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Handler struct {
	App   *echo.Echo
	DB    *gorm.DB
	Viper *viper.Viper
	Log   *logrus.Logger
}

func NewHandler(app *echo.Echo, db *gorm.DB, viper *viper.Viper, log *logrus.Logger) *Handler {
	return &Handler{
		App:   app,
		DB:    db,
		Viper: viper,
		Log:   log,
	}
}
