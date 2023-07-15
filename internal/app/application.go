package app

import entities "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/entity"

type Entities struct {
	UserEntity     *entities.UserEntity
	CategoryEntity *entities.CategoryEntity
	TaskEntity     *entities.TaskEntity
}

func NewEntities() *Entities {
	return &Entities{
		UserEntity:     &entities.UserEntity{},
		CategoryEntity: &entities.CategoryEntity{},
		TaskEntity:     &entities.TaskEntity{},
	}
}
