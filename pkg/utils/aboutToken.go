package utils

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	config "github.com/hafiihzafarhana/DDD-With-Go/internal/configs"
	entities "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/entity"
	repositories "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/repository"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"gorm.io/gorm"
)

var db *gorm.DB = config.GetPostgresInstance()

func parseToken(tokenString string, secretKey string) (*jwt.Token, errors.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewUnauthenticatedError("invalid token error")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.NewUnauthenticatedError(err.Error())
	}

	return token, nil
}

func VerifyToken(bearerToken string) (*entities.UserEntity, errors.MessageErr) {
	var SECRET_ACC_KEY = os.Getenv("SECRET_ACC_KEY")
	bearer := strings.HasPrefix(bearerToken, "Bearer")

	if !bearer || bearerToken == "" {
		return nil, errors.NewUnautorizhedError("Authorization header missing")
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := parseToken(tokenString, SECRET_ACC_KEY)
	if err != nil {
		return nil, err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errors.NewUnauthenticatedError("invalid token error")
	} else {
		mapClaims = claims
	}

	user := entities.UserEntity{
		Email: mapClaims["email"].(string),
		Role:  mapClaims["role"].(string),
	}

	userRepo := repositories.NewUserPG(db)
	getUser, err := userRepo.GetUserByEmail(user)
	if err != nil {
		return nil, err
	}

	return getUser, nil
}
