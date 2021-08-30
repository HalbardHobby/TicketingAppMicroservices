package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	input := data.User{}
	input.FromJson(r.Body)

	found := new(data.User)

	err := data.UserCollection.FindOne(context.TODO(), bson.M{"email": input.Username}).Decode(&found)
	if err != nil {
		be := errors.BadRequestError{
			Reason: "Email and password do not match",
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}

	if !data.CheckPasswordHash(found.Password, input.Password) {
		be := errors.BadRequestError{
			Reason: "Email and password do not match",
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}

	cookie, err := generateAuthenticationCookie(*found)
	if err != nil {
		log.Println(err.Error())
		be := errors.BadRequestError{
			Reason: err.Error(),
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}

	found.Password = ""

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
