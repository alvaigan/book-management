package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;unique;not null" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at,omitzero"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at,omitzero"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Books     []Book         `gorm:"foreignKey:author_id;references:id" json:"books"`
}
