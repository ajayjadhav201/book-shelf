package main

import (
	"log"

	"github.com/ajayjadhav201/book-shelf/database"
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

	routes.RegisterRoutes(r.Group("/v1"))
	//
	//
	log.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
