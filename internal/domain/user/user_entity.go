package user

import (
	"time"

	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/task"
)

type UserEntity struct {
	ID        uint              `gorm:"primaryKey"`
	Full_name string            `gorm:"not null"`
	Email     string            `gorm:"not null;unique;type:varchar(191)"`
	Password  string            `gorm:"not null"`
	Role      string            `gorm:"not null"`
	Tasks     []task.TaskEntity `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
