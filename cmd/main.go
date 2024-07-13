package main

import (
	"log"

	"github.com/ajayjadhav201/book-shelf/database"
	"github.com/ajayjadhav201/book-shelf/middleware"
	"github.com/ajayjadhav201/book-shelf/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db := database.NewDatabase()
	database.SetDatabase(db)
	//
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.Security())
	r.Use(middleware.Xss())
	r.Use(middleware.Cors())
	// r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute
	//
	//
	routes.RegisterRoutes(r.Group("/v1/api"))
	//
	//
	log.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
