package models

import (
	"context"
	"crud-nosql/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Isbn        string `bson:"isbn" json:"isbn"`
	Name        string `bson:"name" json:"name"`
	Publication string `bson:"pub" json:"pub"`
	Author      string `bson:"authorId" json:"authorId"`
}

type BookResponse struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Isbn        string             `bson:"isbn" json:"isbn"`
	Name        string             `bson:"name" json:"name"`
	Publication string             `bson:"pub" json:"pub"`
	Author      string             `bson:"authorId" json:"authorId"`
}

var booksCollection *mongo.Collection
var booksCtx context.Context

func init() {
	db := config.Connect()
	booksCollection = db.Collection("books")
	booksCtx = config.Ctx
}

func (b *Book) CreateBook() (*mongo.InsertOneResult, error) {
	authorId := b.Author
	if _, err := IncrementAuthorBookCount(authorId); err != nil {
		return booksCollection.InsertOne(booksCtx, b)
	} else {
		return nil, err
	}
}

func GetAllBooks() ([]BookResponse, error) {
	var books []BookResponse
	cursor, err := booksCollection.Find(booksCtx, bson.D{})
	if err != nil {
		return books, err
	}
	if err := cursor.All(booksCtx, books); err != nil {
		return books, err
	}
	return books, nil
}

func GetBookById(id string) (BookResponse, error) {
	var book BookResponse
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}

	err = booksCollection.FindOne(booksCtx, bson.M{"_id": obId}).Decode(&book)

	if err != nil {
		return book, err
	}
	return book, nil
}

func UpdateBook(book BookResponse) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: book.Id}}
	update := bson.D{{Key: "$set", Value: book}}
	return booksCollection.UpdateOne(booksCtx, filter, update)
}

func DeleteBookById(id string) (*mongo.DeleteResult, error) {
	var book Book
	obId, _ := primitive.ObjectIDFromHex(id)
	err := booksCollection.FindOne(booksCtx, bson.M{"_id": obId}).Decode(&book)
	if err != nil {
		return nil, err
	}
	if _, err := DecrementAuthorBookCount(book.Author); err != nil {
		return booksCollection.DeleteOne(booksCtx, bson.M{"_id": obId})
	} else {
		return nil, err
	}
}
