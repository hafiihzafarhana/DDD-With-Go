package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/interfaces"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/utils"
)

type UserAuth struct {
	UserID uint
	Email  string
}

func Authenticate(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	getUser, err := utils.VerifyToken(authHeader)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
	}

	userAuth := interfaces.NewJWTAuthenticateTransformer{
		Id:    getUser.ID,
		Email: getUser.Email,
	}

	ctx.Set("user", userAuth)

	ctx.Next()
}
