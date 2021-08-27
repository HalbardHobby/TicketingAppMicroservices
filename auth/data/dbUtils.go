package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(pasword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}

	return string(hash)
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}
	return false
}
