package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx              context.Context
	Client           *mongo.Client
	Db               *mongo.Database
	BookCollection   *mongo.Collection
	AuthorCollection *mongo.Collection
)

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://test:test@cluster0.w7ovegb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	Client = client
	Db = client.Database("go-crud-npsql")
	BookCollection = Db.Collection("books")
	AuthorCollection = Db.Collection("author")
}

func Disconnect() {
	Client.Disconnect(Ctx)
}
