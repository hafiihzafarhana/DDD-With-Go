package task

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      bool   `gorm:"not null"`
	UserID      uint
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
