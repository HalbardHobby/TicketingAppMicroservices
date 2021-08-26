package routes

import (
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := data.User{}

	err := user.FromJson(r.Body)
	if err != nil {
		m := make(map[string]string)
		m["error"] = err.Error()
		errors.JsonError(w, m, http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		// m := make(map[string]string)
		// for _, fieldErr := range err.(validator.ValidationErrors) {
		// 	fmt.Println(fieldErr)
		// 	m[fieldErr.Field()] = fieldErr.Error()
		// 	fmt.Println(m)
		// }

		errors.JsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

}
