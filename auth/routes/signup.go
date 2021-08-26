package routes

import (
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/go-playground/validator/v10"
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

}
