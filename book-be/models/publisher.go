package models

import (
	"time"

	"gorm.io/gorm"
)

type Publisher struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;unique;not null"`
	Name      string         `gorm:"type:varchar(255)"`
	City      string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
