package models

import (
	"context"
	"crud-nosql/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Isbn        string  `bson:"isbn" json:"isbn"`
	Name        string  `bson:"name" json:"name"`
	Publication string  `bson:"pub" json:"pub"`
	Author      *Author `bson:"author" json:"author"`
}

var booksCollection *mongo.Collection
var booksCtx context.Context

func init() {
	db := config.Connect()
	booksCollection = db.Collection("books")
	booksCtx = config.Ctx
}

func (b *Book) CreateBook() *mongo.InsertOneResult {
	result, err := booksCollection.InsertOne(booksCtx, b)
	if err != nil {
		panic(err)
	}
	return result
}

func GetAllBooks() []Book {
	var books []Book
	cursor, err := booksCollection.Find(booksCtx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err := cursor.All(booksCtx, books); err != nil {
		panic(err)
	}
	return books
}

func GetBookById(id string) Book {
	var book Book
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	err = booksCollection.FindOne(booksCtx, bson.M{"_id": obId}).Decode(book)

	if err != nil {
		panic(err)
	}
	return book
}

func UpdateBook(id string, book Book) *mongo.UpdateResult {
	obId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: book}}
	result, err := booksCollection.UpdateOne(booksCtx, filter, update)
	if err != nil {
		panic(err)
	}
	return result
}

func DeleteBookById(id string) *mongo.DeleteResult {
	obId, _ := primitive.ObjectIDFromHex(id)
	result, err := booksCollection.DeleteOne(booksCtx, bson.M{"_id": obId})
	if err != nil {
		panic(err)
	}
	return result
}
