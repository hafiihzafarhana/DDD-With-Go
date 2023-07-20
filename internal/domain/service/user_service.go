package services

import (
	repositories "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/repository"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/dtos"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/interfaces"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
)

type UserService interface {
	Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr)
	Login(payload dtos.NewLoginRequest) (*interfaces.NewLoginResponse, errors.MessageErr)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload dtos.NewRegisterRequest) (*interfaces.NewRegisterResponse, errors.MessageErr) {
	user := payload.RegisterRequestToEntity()

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

func (u *userService) Login(payload dtos.NewLoginRequest) (*interfaces.NewLoginResponse, errors.MessageErr) {
	user := payload.LoginRequestToEntity()

	getUser, err := u.userRepo.GetUserByEmail(user)
	if err != nil {
		return nil, err
	}

	if err := getUser.ComparePassword(user.Password); err != nil {
		return nil, err
	}

	refreshToken, accToken := getUser.GenerateToken()

	response := &interfaces.NewLoginResponse{
		RefreshToken: refreshToken,
		AccToken:     accToken,
	}

	return response, nil
}
