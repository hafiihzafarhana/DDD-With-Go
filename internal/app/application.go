package app

import (
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/category"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/task"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/user"
)

type Entities struct {
	UserEntity     *user.UserEntity
	CategoryEntity *category.CategoryEntity
	TaskEntity     *task.TaskEntity
}

func NewEntities() *Entities {
	return &Entities{
		UserEntity:     &user.UserEntity{},
		CategoryEntity: &category.CategoryEntity{},
		TaskEntity:     &task.TaskEntity{},
	}
}
