package seeders

import (
	"book-be/models"
	"book-be/utils"
)

func (s *Seeder) UserSeed() {

	password, err := utils.HashPassword("password")
	if err != nil {
		s.Log.Fatal("Error hashing password")
	}

	if err = s.DB.Create(&models.User{
		Username: "admin@mail.com",
		Password: password,
	}).Error; err != nil {
		s.Log.Error(err.Error())
	}
}
