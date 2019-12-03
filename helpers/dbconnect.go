package helpers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client -- mongo client
var Client *mongo.Client
var once sync.Once
var err error

//ConnectToMongoDb -- coonection to the mongodb database
func ConnectToMongoDb() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017,localhost:27018,localhost:27019/northwind?replicaSet=rs")
		Client, err = mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}
		err = Client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB!")
	})
	return Client
}
