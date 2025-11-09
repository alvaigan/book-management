package models

import (
	"time"

	"gorm.io/gorm"
)

type Publisher struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;unique;not null" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	City      string         `gorm:"type:varchar(255);not null" json:"city"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at,omitzero"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at,omitzero"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Book      *Book          `gorm:"foreignKey:publisher_id;references:id" json:"book,omitempty"`
}
