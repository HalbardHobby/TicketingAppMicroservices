package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	user.FromJson(r.Body)

	count, _ := data.UserCollection.CountDocuments(context.TODO(), bson.M{"email": user.Username})
	if count > 0 {
		errMessage := fmt.Sprintf("Email '%s' already in use", user.Username)
		errors.BadRequestError(w, errMessage)
		return
	}

	user.Password = data.HashPassword(user.Password)
	user.Id = primitive.NewObjectID()
	data.UserCollection.InsertOne(context.TODO(), user)
	user.Password = ""
	log.Printf("Created new User with id: %s and Email: %s", user.Id.Hex(), user.Username)

	cookie, err := generateAuthenticationCookie(user)
	if err != nil {
		errors.BadRequestError(w, err.Error())
		return
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
