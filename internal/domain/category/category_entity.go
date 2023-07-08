package category

import (
	"time"

	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/task"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	Tasks     []task.Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
