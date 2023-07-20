package repositories

import (
	"fmt"

	entities "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/entity"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Register(payload UserEntity)
	Register(payload entities.UserEntity) (*entities.UserEntity, errors.MessageErr)
	GetUserByEmail(payload entities.UserEntity) (*entities.UserEntity, errors.MessageErr)
}

// ini harus lokal
type userRepository struct {
	db *gorm.DB
}

// return harus sesuai dengan UserRepository
func NewUserPG(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Register(payload entities.UserEntity) (*entities.UserEntity, errors.MessageErr) {
	if err := u.db.Create(&payload).Error; err != nil {
		return nil, errors.NewConflictError(err.Error())
	}

	return &payload, nil
}

func (u *userRepository) GetUserByEmail(payload entities.UserEntity) (*entities.UserEntity, errors.MessageErr) {
	var user entities.UserEntity

	if err := u.db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		fmt.Println(err)
		return nil, errors.NewNotFoundError(fmt.Sprintf("user with email %s is not found", payload.Email))
	}

	return &user, nil
}
