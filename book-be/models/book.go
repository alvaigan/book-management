package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;unique;not null" json:"id"`
	Title       string         `gorm:"type:tinytext;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	AuthorId    uint           `gorm:"not null" json:"author_id"`
	PublisherId uint           `gorm:"not null" json:"publisher_id"`
	CreatedAt   time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
