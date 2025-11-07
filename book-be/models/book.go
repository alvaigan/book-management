package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;unique;not null"`
	Title       string         `gorm:"type:tinytext;not null"`
	Description string         `gorm:"type:text"`
	PublisherID uint           `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
