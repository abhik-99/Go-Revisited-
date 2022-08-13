package config

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx    context.Context
	Client *mongo.Client
)

func Connect() mongo.Database {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("../../.env")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.w7ovegb.mongodb.net/?retryWrites=true&w=majority", myEnv["DBUser"], myEnv["DBPass"])))
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	Client = client
	return *client.Database("go-crud-npsql")
}

func Disconnect() {
	Client.Disconnect(Ctx)
}
