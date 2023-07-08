package route

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	config "github.com/hafiihzafarhana/DDD-With-Go/internal/configs"
	routes "github.com/hafiihzafarhana/DDD-With-Go/internal/interfaces/api/routes/routes"
)

var PORT string

func StartApp() {
	config.GetPostgresInstance()

	r := gin.Default()
	route := r.Group("/api")

	routes.UserRoute(route)

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}
	log.Fatalln(r.Run(":" + PORT))
}
