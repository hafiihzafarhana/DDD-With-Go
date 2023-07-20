package entities

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID        uint         `gorm:"primaryKey"`
	Full_name string       `gorm:"not null"`
	Email     string       `gorm:"not null;unique;type:varchar(191)"`
	Password  string       `gorm:"not null"`
	Role      string       `gorm:"not null"`
	Tasks     []TaskEntity `gorm:"foreignKey:UserID"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
}

func (u *UserEntity) HashPassword() errors.MessageErr {
	const cost = 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	u.Password = string(bs)
	return nil
}

func (u *UserEntity) ComparePassword(reqPassword string) errors.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(reqPassword)); err != nil {
		return errors.NewBadRequestError("wrong password")
	}

	return nil
}

func (u *UserEntity) tokenClaims() (jwt.MapClaims, jwt.MapClaims) {
	expirationRefreshTime := time.Now().Add(time.Hour * 24)
	expirationAccTime := time.Now().Add(time.Minute)
	issuedAtTime := time.Now()
	return jwt.MapClaims{
			"id":    u.ID,
			"email": u.Email,
			"role":  u.Role,
			"iss":   "ddd",                        //menunjukan token yang diterbitkan oleh pemilik apk
			"sub":   u.ID,                         // menunjukan user yang telah melakukan clain token
			"exp":   expirationRefreshTime.Unix(), // menunjukan waktu berakhir
			"iat":   issuedAtTime.Unix(),          // menunjukan waktu pertama claim
		}, jwt.MapClaims{
			"id":    u.ID,
			"email": u.Email,
			"role":  u.Role,
			"iss":   "ddd",                    //menunjukan token yang diterbitkan oleh pemilik apk
			"sub":   u.ID,                     // menunjukan user yang telah melakukan clain token
			"exp":   expirationAccTime.Unix(), // menunjukan waktu berakhir
			"iat":   issuedAtTime.Unix(),      // menunjukan waktu pertama claim
		}
}

func (u *UserEntity) signToken(claims jwt.MapClaims, secretKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken
}

func (u *UserEntity) GenerateToken() (string, string) {
	secretRefreshKey := os.Getenv("SECRET_REFRESH_KEY")
	secretAccKey := os.Getenv("SECRET_ACC_KEY")
	refresh, acc := u.tokenClaims()

	return u.signToken(refresh, secretRefreshKey), u.signToken(acc, secretAccKey)
}
