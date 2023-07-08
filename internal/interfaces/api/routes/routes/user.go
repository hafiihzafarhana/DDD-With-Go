package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	routeGroup := r.Group("/users")

	routeGroup.POST("/register")
	routeGroup.POST("/login")
}
