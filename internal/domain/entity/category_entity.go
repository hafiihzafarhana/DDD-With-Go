package entities

import (
	"time"
)

type CategoryEntity struct {
	ID        uint         `gorm:"primaryKey"`
	Type      string       `gorm:"not null"`
	Tasks     []TaskEntity `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
