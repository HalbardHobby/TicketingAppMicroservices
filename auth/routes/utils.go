package routes

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

func ValidateUserInputMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Copying data to allow second read by handler
		raw, _ := ioutil.ReadAll(r.Body)
		reader := bytes.NewReader(raw)

		user := data.User{}
		err := user.FromJson(reader)
		if err != nil {
			je := errors.JsonFormattingError{
				Reason: err.Error(),
				Code:   http.StatusBadRequest}
			errors.JsonError(rw, &je)
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

func generateAuthenticationCookie(user data.User) (*http.Cookie, error) {
	// Generate JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}
	// Save on session
	return &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	}, nil

}
