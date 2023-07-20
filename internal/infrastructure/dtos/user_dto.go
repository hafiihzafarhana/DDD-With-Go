package dtos

import (
	"time"

	entities "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/entity"
)

type NewRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type NewLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *NewRegisterRequest) RegisterRequestToEntity() entities.UserEntity {
	return entities.UserEntity{
		Full_name: u.FullName,
		Email:     u.Email,
		Password:  u.Password,
		Role:      "member",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *NewLoginRequest) LoginRequestToEntity() entities.UserEntity {
	return entities.UserEntity{
		Email:    u.Email,
		Password: u.Password,
	}
}
