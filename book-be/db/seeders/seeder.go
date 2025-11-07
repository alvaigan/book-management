package seeders

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Seeder struct {
	Log *logrus.Logger
	DB  *gorm.DB
}
