package main

import (
	"book-be/config"
	"book-be/models"
	"fmt"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogrus()
	db := config.NewDatabase(viper, log)

	m := gormigrate.New(db.Debug().Begin(), gormigrate.DefaultOptions, []*gormigrate.Migration{})

	argsRaw := os.Args
	args := argsRaw[1:]

	if len(args) == 0 {
		fmt.Println("No migration executed!")
	} else {
		cmd := args[0]

		switch cmd {
		case "migrate":
			if len(args) > 1 && args[1] == "fresh" {
				m.InitSchema(func(tx *gorm.DB) error {
					err := tx.AutoMigrate(
						&models.User{},
						&models.Author{},
						&models.Book{},
						&models.Publisher{},
					)
					if err != nil {
						return err
					}
					return nil
				})
			}

			if err := m.Migrate(); err != nil {
				log.Fatalf("Migration failed: %v", err)
			}
			fmt.Println("Migration did run successfully")
		case "rollback":
			if err := m.RollbackLast(); err != nil {
				log.Fatalf("Rolling bank failed: %v", err)
			}
			fmt.Println("Rolling back did run successfully")
		}

	}
}
