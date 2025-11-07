package routes

import (
	"book-be/utils"

	"github.com/labstack/echo/v4"
)

func (rc *RouteConfig) GuestRoute() {
	api := rc.App
	api.GET("/", func(c echo.Context) error {
		return c.JSON(200, utils.GenerateRes("Book Management API", nil))
	})
}
