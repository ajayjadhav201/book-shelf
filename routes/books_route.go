package routes

import (
	"github.com/ajayjadhav201/book-shelf/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	accounts := gin.Accounts{
		"admin": "admin",
		"test":  "test",
	}
	//
	r.Use(gin.BasicAuth(accounts))
	{
		r.GET("/books", controllers.GetBooksHandler)
		r.GET("/books/:id", controllers.GetBookByIdHandler)
		r.POST("/books", controllers.InsertBookHandler)
		r.PATCH("/books", controllers.UpdateBookHandler)
		r.DELETE("/books/:id", controllers.DeleteBookHandler)
		r.GET("/genres", controllers.GetBookByGenreHandler)
	}
}
