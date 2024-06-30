package controllers

import (
	"github.com/ajayjadhav201/book-shelf/database"
	"github.com/ajayjadhav201/book-shelf/utils"
	"github.com/gin-gonic/gin"
)

var (
	Ok                        = 200
	StatusInternalServerError = 500
)

func GetBooksHandler(c *gin.Context) {
	//
	books, err := database.GetAllBooks(1, 1)
	if err != nil {
		utils.WriteInternalServerError(c)
	}
	utils.WriteJson(c, books)
}

func GetBookByIdHandler(c *gin.Context) {
	//
}

func GetBookByGenreHandler(c *gin.Context) {
	//
}

func InsertBookHandler(c *gin.Context) {
	//
}

func UpdateBookHandler(c *gin.Context) {
	//
}

func DeleteBookHandler(c *gin.Context) {
	//
}
