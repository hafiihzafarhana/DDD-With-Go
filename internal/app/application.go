package app

import (
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/category"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/task"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/user"
)

type Entities struct {
	User     *user.User
	Category *category.Category
	Task     *task.Task
}

func NewEntities() *Entities {
	return &Entities{
		User:     &user.User{},
		Category: &category.Category{},
		Task:     &task.Task{},
	}
}
