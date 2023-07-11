package user

import (
	"fmt"

	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/dtos"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/interfaces"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
)

type UserService interface {
	Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr) {
	user := payload
	fmt.Println(user)
	return nil, nil
	// err := user.HashPassword()
	// if err != nil {
	// 	return nil, err
	// }

	// createdUser, err := u.userRepo.Register(user)
	// if err != nil {
	// 	return nil, err
	// }

	// response := &interfaces.NewRegisterResponse{
	// 	Id:        createdUser.ID,
	// 	FullName:  createdUser.Full_name,
	// 	Email:     createdUser.Email,
	// 	CreatedAt: createdUser.CreatedAt,
	// }

	// return response, nil
}
