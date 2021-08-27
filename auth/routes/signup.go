package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := data.User{}

	err := user.FromJson(r.Body)
	if err != nil {
		je := new(errors.JsonFormattingError)
		je.Reason = err.Error()
		errors.JsonError(w, je, http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		ve := new(errors.RequestValidationError)
		ve.Reasons = err.(validator.ValidationErrors)
		errors.JsonError(w, ve, http.StatusBadRequest)
		return
	}

	count, _ := data.UserCollection.CountDocuments(context.TODO(), bson.M{"email": user.Username})
	if count > 0 {
		log.Printf("Email '%s' already in use", user.Username)
		return
	}

	data.UserCollection.InsertOne(context.TODO(), user)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
