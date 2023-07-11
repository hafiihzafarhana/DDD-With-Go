package route

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	config "github.com/hafiihzafarhana/DDD-With-Go/internal/configs"
	"github.com/hafiihzafarhana/DDD-With-Go/internal/infrastructure/api/routes/routes"
)

var PORT string

func StartApp() {
	db := config.GetPostgresInstance()

	r := gin.Default()
	route := r.Group("/api")

	routes.UserRoute(route, db)

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}
	log.Fatalln(r.Run(":" + PORT))
}
