package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := data.User{}

	err := user.FromJson(r.Body)
	if err != nil {
		je := errors.JsonFormattingError{
			Reason: err.Error(),
			Code:   http.StatusBadRequest}
		errors.JsonError(w, &je)
		return
	}

	err = user.Validate()
	if err != nil {
		ve := errors.RequestValidationError{
			Reasons: err.(validator.ValidationErrors),
			Code:    http.StatusBadRequest}
		errors.JsonError(w, &ve)
		return
	}

	count, _ := data.UserCollection.CountDocuments(context.TODO(), bson.M{"email": user.Username})
	if count > 0 {
		errMessage := fmt.Sprintf("Email '%s' already in use", user.Username)
		log.Println(errMessage)
		be := errors.BadRequestError{
			Reason: errMessage,
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}

	user.Password = data.HashPassword(user.Password)
	user.Id = primitive.NewObjectID()
	data.UserCollection.InsertOne(context.TODO(), user)
	log.Printf("Created new User with id: %s and Email: %s", user.Id.Hex(), user.Username)

	// Generate JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := at.SignedString([]byte("asdf"))
	if err != nil {
		log.Println(err.Error())
		be := errors.BadRequestError{
			Reason: err.Error(),
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}
	// Save on session
	cookie := &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
