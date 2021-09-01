package routes

import (
	"context"
	"encoding/json"
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
		errors.BadRequestError(w, "Email and password do not match")
		return
	}

	if !data.CheckPasswordHash(found.Password, input.Password) {
		errors.BadRequestError(w, "Email and password do not match")
		return
	}

	cookie, err := generateAuthenticationCookie(*found)
	if err != nil {
		errors.BadRequestError(w, "Email and password do not match")
		return
	}

	found.Password = ""

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(found)
}
