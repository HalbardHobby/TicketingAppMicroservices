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
	"go.mongodb.org/mongo-driver/bson"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
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
	if count == 0 {
		errMessage := fmt.Sprintf("Email '%s' not found", user.Username)
		log.Println(errMessage)
		be := errors.BadRequestError{
			Reason: errMessage,
			Code:   404,
		}
		errors.JsonError(w, &be)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
