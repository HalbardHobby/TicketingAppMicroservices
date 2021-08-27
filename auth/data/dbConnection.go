package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://auth-mongo-service:27017")
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to DB")
	}
	Client = c
	UserCollection = Client.Database("Auth").Collection("Users")
}
