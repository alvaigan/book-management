package main

import (
	"book-be/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogrus()
	db := config.NewDatabase(viper, log)
	app := echo.New()

	config.NewApp(&config.AppConfig{
		App:   app,
		DB:    db,
		Viper: viper,
		Log:   log,
	})

	fmt.Println(viper.GetString("app.port"))

	err := app.Start(":" + viper.GetString("app.port"))
	if err != nil {
		log.Panic(err)
	}
}
