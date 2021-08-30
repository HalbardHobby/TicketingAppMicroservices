package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/golang-jwt/jwt"
)

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("session")
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"currentUser": "null"})
	}

	payload, err := jwt.ParseWithClaims(tokenCookie.Value, new(data.User), func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		errMessage := err.Error()
		be := errors.BadRequestError{
			Reason: errMessage,
			Code:   400,
		}
		errors.JsonError(w, &be)
		return
	}

	if !payload.Valid {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"currentUser": "null"})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(payload.Claims)
}
