package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/domain/user"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/api/handlers"
	"gorm.io/gorm"
)

func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := user.NewUserPG(db)
	userService := user.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routeGroup := r.Group("/users")
	routeGroup.POST("/register", userHandler.Register)
}
