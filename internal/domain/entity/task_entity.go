package entities

import "time"

type TaskEntity struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      bool   `gorm:"not null"`
	UserID      uint   `gorm:"index"`
	CategoryID  uint   `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
