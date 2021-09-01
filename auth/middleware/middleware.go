package middleware

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

const UserContextKey = "User"

func GetCurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("session")
		if err != nil {
			next.ServeHTTP(rw, r)
		}

		payload, err := jwt.ParseWithClaims(tokenCookie.Value, new(data.User), func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			next.ServeHTTP(rw, r)
		}

		ctx := context.WithValue(r.Context(), UserContextKey, payload.Claims)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func ValidateUserInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Copying data to allow second read by handler
		raw, _ := ioutil.ReadAll(r.Body)
		reader := bytes.NewReader(raw)

		user := data.User{}
		err := user.FromJson(reader)
		if err != nil {
			errors.JsonFormattingError(rw, err.Error())
			return
		}

		err = user.Validate()
		if err != nil {
			ve := errors.RequestValidationError{
				Reasons: err.(validator.ValidationErrors),
				Code:    http.StatusBadRequest}
			errors.JsonError(rw, &ve)
			return
		}

		// Reproducing IO reader for request body
		r.Body = ioutil.NopCloser(bytes.NewBuffer(raw))
		next.ServeHTTP(rw, r)
	})
}
