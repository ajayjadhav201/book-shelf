package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BookID         string
	Title          string
	Author         []string
	Language       string
	Publisher      string
	PublishingDate string
	Genre          string
	ISBN           string
	Edition        string
	Pages          string
	Summery        string
	AboutAuthor    []string
}

func NewBook() {
	//
}
