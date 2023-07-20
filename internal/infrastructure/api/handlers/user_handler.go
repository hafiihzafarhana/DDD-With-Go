package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	services "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/service"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/dtos"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/errors"
	"github.com/hafiihzafarhana/DDD-With-Go/pkg/success"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dtos.NewRegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		errCasting, ok := err.(validator.ValidationErrors)
		// apabila data json tidak ada
		if !ok {
			newErrBind := errors.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		// apabila ada data json, tetapi data tersebut tidak sesuai dengan NewRegisterRequest
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

	ctx.JSON(http.StatusCreated, success.NewCreatedSuccess("Register Success", createdUser))
}

func (u *userHandler) Login(ctx *gin.Context) {
	var requestBody dtos.NewLoginRequest

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

	token, errResponse := u.userService.Login(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, success.NewCreatedSuccess("Login Success", token))
}

func (u *userHandler) GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.MustGet("user"))
}
