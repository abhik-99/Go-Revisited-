package models

import (
	"context"
	"crud-nosql/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Author struct {
	Firstname    string `bson:"firstname" json:"firstname"`
	Lastname     string `bson:"lastname" json:"lastname"`
	BooksWritten uint   `bson:"numBooks" json:"numBooks"`
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

func (a *Author) CreateAuthor() *mongo.InsertOneResult {
	result, err := authorsCollection.InsertOne(authorsCtx, a)
	if err != nil {
		panic(err)
	}
	return result
}

func GetAllAuthors() []Author {
	var authors []Author
	cursor, err := authorsCollection.Find(authorsCtx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err := cursor.All(authorsCtx, authors); err != nil {
		panic(err)
	}
	return authors
}

func GetAuthorById(id string) Author {
	var author Author
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	err = authorsCollection.FindOne(authorsCtx, bson.M{"_id": obId}).Decode(author)

	if err != nil {
		panic(err)
	}
	return author
}

func UpdateAuthor(id string, author Author) *mongo.UpdateResult {
	obId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: author}}
	result, err := authorsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return result
}

func DeleteAuthorById(id string) *mongo.DeleteResult {
	obId, _ := primitive.ObjectIDFromHex(id)
	result, err := authorsCollection.DeleteOne(authorsCtx, bson.M{"_id": obId})
	if err != nil {
		panic(err)
	}
	return result
}
