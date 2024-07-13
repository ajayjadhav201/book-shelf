package routes

import (
	controller "github.com/ajayjadhav201/book-shelf/controllers"
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
		r.GET("/books", controller.GetBooksHandler)
		r.GET("/books/:id", controller.GetBookByIdHandler)
		r.POST("/books", controller.InsertBookHandler)
		r.PUT("/books", controller.UpdateBookHandler)
		r.DELETE("/books/:id", controller.DeleteBookHandler)
		r.GET("/genres/:genre", controller.GetBookByGenreHandler)
		//
		// r.GET("/user", controller.)
	}
}
