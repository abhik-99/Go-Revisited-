package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Db  *mongo.Database
	Ctx context.Context

	client *mongo.Client
	cancel context.CancelFunc
)

func Connect() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("../../.env")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.w7ovegb.mongodb.net/?retryWrites=true&w=majority", myEnv["DBUser"], myEnv["DBPass"])).
		SetServerAPIOptions(serverAPIOptions)
	Ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(Ctx, clientOptions)
	Db = client.Database("crud-auth-nosql")
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	cancel()
	client.Disconnect(Ctx)
}
