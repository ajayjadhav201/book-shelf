package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ajayjadhav201/book-shelf/database"
	model "github.com/ajayjadhav201/book-shelf/models"
	"github.com/ajayjadhav201/book-shelf/utils"
	"github.com/gin-gonic/gin"
)

func GetBooksHandler(c *gin.Context) {
	//
	var pageno int64 = 1
	var limitno int64 = 10
	var err error = nil
	//
	if page := c.Query("page"); page != "" {
		pageno, err = strconv.ParseInt(page, 10, 0)
		if err != nil {
			utils.WriteError(c, http.StatusBadRequest, "Invalid page number")
			return
		}
	}
	//
	if limit := c.Query("limit"); limit != "" {
		limitno, err = strconv.ParseInt(limit, 10, 0)
		if err != nil {
			utils.WriteError(c, http.StatusBadRequest, "Invalid limit number")
			return
		}
	}
	//
	log.Println("request with page: ", pageno, " and limit: ", limitno)
	books, err := database.GetAllBooks(pageno, limitno, map[string]interface{}{})
	if err != nil {
		response := fmt.Sprintf("Unable to get book: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
	utils.WriteJson(c, books)
}

func GetBookByIdHandler(c *gin.Context) {
	//
	id, ok := c.Params.Get("id")
	if !ok {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book id")
	}
	//
	books, err := database.GetBookById(id)
	if err != nil {
		response := fmt.Sprintf("Unable to get book: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
	utils.WriteJson(c, books)
}

func GetBookByGenreHandler(c *gin.Context) {
	//
	genre, ok := c.Params.Get("genre")
	if !ok {
		utils.WriteError(c, http.StatusBadRequest, "Invalid genre")
	}
	filter := make(map[string]interface{})
	filter["genre"] = genre
	//
	books, err := database.GetAllBooks(1, 1, filter)
	if err != nil {
		response := fmt.Sprintf("Unable to get books: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
	utils.WriteJson(c, books)
}

func InsertBookHandler(c *gin.Context) {
	var book *model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book object")
		return
	}
	created := time.Now()
	updated := time.Now()
	book.CreatedAt = &created
	book.UpdatedAt = &updated
	//
	err = database.InsertBook(book)
	if err != nil {
		response := fmt.Sprintf("Unable to insert book: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
	utils.WriteJson(c, "Book inserted successfully")
}

func UpdateBookHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book id")
	}
	//
	var book *model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book object")
		return
	}
	jsonString, err := json.Marshal(book)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book object")
		return
	}
	var updateData map[string]interface{}
	// Unmarshal the JSON into the map
	err = json.Unmarshal([]byte(jsonString), &updateData)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book object")
		return
	}
	err = database.UpdateBook(id, updateData)
	if err != nil {
		response := fmt.Sprintf("Unable to update book: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
}

func DeleteBookHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		utils.WriteError(c, http.StatusBadRequest, "Invalid book id")
	}
	err := database.DeleteBook(id)
	if err != nil {
		response := fmt.Sprintf("Unable to delete book: %s", err.Error())
		utils.WriteError(c, http.StatusInternalServerError, response)
		return
	}
}
