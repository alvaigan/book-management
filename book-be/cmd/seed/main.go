package main

import (
	"book-be/config"
	"book-be/db/seeders"
	"fmt"
	"os"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogrus()
	db := config.NewDatabase(viper, log)

	seeder := seeders.Seeder{
		DB:  db,
		Log: log,
	}

	argsRaw := os.Args
	args := argsRaw[1:]

	if len(args) == 0 {
		fmt.Println("No seed executed!")
	} else {
		cmd := args[0]

		switch cmd {
		case "user_seeder":
			seeder.UserSeed()
		}

		fmt.Println("Seed did run successfully")
	}
}
