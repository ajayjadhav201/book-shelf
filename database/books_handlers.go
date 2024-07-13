package database

import (
	"context"
	"time"

	model "github.com/ajayjadhav201/book-shelf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *Database

func SetDatabase(database *Database) {
	db = database
}

func GetAllBooks(page int64, limit int64, filter interface{}) ([]*model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	books := []*model.Book{}
	//
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := db.Books.Find(ctx, filter, findOptions)

	if err != nil {
		return books, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book model.Book
		if err = cursor.Decode(&book); err != nil {
			return nil, err
		}
		//
		books = append(books, &book)
	}
	return books, nil
	//
}

func GetBookById(id string) ([]*model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	books := []*model.Book{}
	//
	findOptions := options.Find()
	filter := make(map[string]interface{})
	filter["book_id"] = id
	//
	cursor, err := db.Books.Find(ctx, filter, findOptions)

	if err != nil {
		return books, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book model.Book
		if err = cursor.Decode(&book); err != nil {
			return nil, err
		}
		//
		books = append(books, &book)
	}
	return books, nil
	//
}

func InsertBook(book *model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Books.InsertOne(ctx, book)
	return err
}

func UpdateBook(id string, updateData map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	filter := bson.M{"BookID": id}

	update := bson.M{
		"$set": updateData,
	}
	//
	_, err := db.Books.UpdateOne(ctx, filter, update)
	return err
}

func DeleteBook(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//
	filter := bson.M{"BookID": id}
	_, err := db.Books.DeleteOne(ctx, filter)
	return err
}
