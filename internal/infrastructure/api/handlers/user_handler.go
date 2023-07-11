package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/user"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/dtos"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) userHandler {
	return userHandler{userService: userService}
}

func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dtos.NewRegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		errCasting, ok := err.(validator.ValidationErrors)
		if !ok {
			newErrBind := errors.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errors.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	createdUser, errResponse := u.userService.Register(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}