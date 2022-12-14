package models

import (
	"context"
	"crud-nosql/pkg/config"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Author struct {
	Firstname    string `bson:"firstname" json:"firstname"`
	Lastname     string `bson:"lastname" json:"lastname"`
	BooksWritten uint   `bson:"numBooks,omitempty" json:"numBooks"`
}

type AuthorResponse struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	Firstname    string             `bson:"firstname" json:"firstname"`
	Lastname     string             `bson:"lastname" json:"lastname"`
	BooksWritten uint               `bson:"numBooks,omitempty" json:"numBooks"`
}

var (
	authorsCollection *mongo.Collection
	authorsCtx        context.Context
)

func init() {
	db := config.Connect()
	authorsCollection = db.Collection("authors")
	authorsCtx = config.Ctx
}

func (a *Author) CreateAuthor() (*mongo.InsertOneResult, error) {
	a.BooksWritten = 0
	return authorsCollection.InsertOne(authorsCtx, *a)
}

func GetAllAuthors() ([]AuthorResponse, error) {
	var authors []AuthorResponse
	cursor, err := authorsCollection.Find(authorsCtx, bson.D{})
	if err != nil {
		return authors, err
	}
	if err := cursor.All(authorsCtx, &authors); err != nil {
		return authors, err
	}
	if len(authors) == 0 {
		return authors, fmt.Errorf("No Authors defined!")
	}
	return authors, nil
}

func GetAuthorById(id string) (AuthorResponse, error) {
	var author AuthorResponse
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return author, err
	}

	err = authorsCollection.FindOne(authorsCtx, bson.M{"_id": obId}).Decode(&author)

	if err != nil {
		return author, err
	}
	return author, nil
}

func UpdateAuthor(author AuthorResponse) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: author.Id}}
	update := bson.D{{Key: "$set", Value: author}}

	return authorsCollection.UpdateOne(authorsCtx, filter, update)
}

func DeleteAuthorById(id string) (*mongo.DeleteResult, error) {
	var author Author
	obId, _ := primitive.ObjectIDFromHex(id)
	err := authorsCollection.FindOne(authorsCtx, bson.M{"_id": obId}).Decode(&author)
	if err != nil {
		return nil, err
	}
	if author.BooksWritten != 0 {
		return nil, fmt.Errorf("Author has atleast 1 book in store!")
	}
	return authorsCollection.DeleteOne(authorsCtx, bson.M{"_id": obId})
}

func IncrementAuthorBookCount(id string) (*mongo.UpdateResult, error) {
	var author Author
	obId, _ := primitive.ObjectIDFromHex(id)
	err := authorsCollection.FindOne(authorsCtx, bson.M{"_id": obId}).Decode(author)
	if err != nil {
		return nil, err
	}
	author.BooksWritten += 1

	filter := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: author}}

	return authorsCollection.UpdateOne(authorsCtx, filter, update)
}

func DecrementAuthorBookCount(id string) (*mongo.UpdateResult, error) {
	var author Author
	obId, _ := primitive.ObjectIDFromHex(id)
	err := authorsCollection.FindOne(authorsCtx, bson.M{"_id": obId}).Decode(author)
	if err != nil {
		return nil, err
	}
	author.BooksWritten -= 1

	filter := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: author}}

	return authorsCollection.UpdateOne(authorsCtx, filter, update)
}
