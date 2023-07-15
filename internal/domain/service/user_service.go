package services

import (
	"fmt"

	repositories "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/repository"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/dtos"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/interfaces"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
)

type UserService interface {
	Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr) {
	user := payload.RegisterRequestToEntity()
	fmt.Println(user)
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	createdUser, err := u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &interfaces.NewRegisterResponse{
		Id:        createdUser.ID,
		FullName:  createdUser.Full_name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}

	return response, nil
}
