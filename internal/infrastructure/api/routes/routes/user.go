package routes

import (
	"github.com/gin-gonic/gin"
	repositories "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/repository"
	services "github.com/hafiihzafarhana/DDD-With-Go/internal/domain/service"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/api/handlers"
	"gorm.io/gorm"
)

func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routeGroup := r.Group("/users")
	routeGroup.POST("/register", userHandler.Register)
}
