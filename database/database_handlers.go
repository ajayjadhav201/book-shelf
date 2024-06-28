package database

import (
	"context"
	"time"

	"github.com/ajayjadhav201/book-shelf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *Database

func SetDatabase(database *Database) {
	db = database
}

func GetAllBooks(page int64, limit int64) ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//
	books := []*models.Book{}
	//
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := db.Collection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return books, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		if err = cursor.Decode(&book); err != nil {
			return nil, err
		}
		//
		books = append(books, &book)
	}
	return books, nil
	//
}

func InsertBook(book *models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := db.Collection.InsertOne(ctx, book)
	return err
}

func UpdateBook(id string, book *models.Book) {
	//
}
