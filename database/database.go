package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ajayjadhav201/book-shelf/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewDatabase() *Database {
	client := dbinit()
	coll := openCollection(client, "books")
	return &Database{Client: client, Collection: coll}
}

func dbinit() *mongo.Client {
	MongoDb := utils.GetEnv("MONGO_URL", "")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal("Error while connecting to mongodb: ", err)
	}
	//
	fmt.Println("Connected to MongoDB")
	return client
}

func openCollection(client *mongo.Client, CollectionName string) *mongo.Collection {
	fmt.Println("Opening Collection: ", CollectionName)
	var collection *mongo.Collection = client.Database("book_shelf").Collection(CollectionName)
	return collection
}

//
