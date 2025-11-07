package seeders

import (
	"book-be/models"
)

func (s *Seeder) DummyData() {
	publisher := models.Publisher{
		Name: "Gramedia Publisher",
		City: "Jakarta",
	}
	author := models.Author{
		Name: "Paulo Coelho",
	}

	if err := s.DB.Create(&publisher).Error; err != nil {
		s.Log.Error(err.Error())
	}

	if err := s.DB.Create(&author).Error; err != nil {
		s.Log.Error(err.Error())
	}

	if err := s.DB.Create(&models.Book{
		Title:       "The Alchemist",
		Description: "The Journey of Alchemist",
		PublisherId: publisher.ID,
		AuthorId:    author.ID,
	}).Error; err != nil {
		s.Log.Error(err.Error())
	}
}
