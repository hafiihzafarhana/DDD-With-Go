package repositories

import (
	entities "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/entity"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Register(payload UserEntity)
	Register(payload entities.UserEntity) (*entities.UserEntity, errors.MessageErr)
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
		return nil, errors.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}
