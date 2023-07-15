package entities

import (
	"time"

	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID        uint              `gorm:"primaryKey"`
	Full_name string            `gorm:"not null"`
	Email     string            `gorm:"not null;unique;type:varchar(191)"`
	Password  string            `gorm:"not null"`
	Role      string            `gorm:"not null"`
	Tasks     []TaskEntity `gorm:"foreignKey:UserID"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime"`
}

func (u *UserEntity) HashPassword() errors.MessageErr {
	const cost = 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return errors.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)
	return nil
}