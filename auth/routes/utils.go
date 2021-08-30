package routes

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
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
