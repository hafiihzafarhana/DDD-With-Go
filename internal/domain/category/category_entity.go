package category

import (
	"time"

	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/task"
)

type CategoryEntity struct {
	ID        uint              `gorm:"primaryKey"`
	Type      string            `gorm:"not null"`
	Tasks     []task.TaskEntity `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
