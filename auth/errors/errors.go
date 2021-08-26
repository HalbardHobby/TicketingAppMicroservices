package errors

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type RequestValidationError struct {
	Reasons []validator.FieldError
}

func (e *RequestValidationError) Error() string {
	var err strings.Builder
	for _, e := range e.Reasons {
		err.WriteString(e.Error() + "\n")
	}
	return err.String()
}

type JsonFormattingError struct {
	Reason string
}

func (e *JsonFormattingError) Error() string {
	return e.Reason
}

type errorResponse struct {
	Errors []errorMessage `json:"errors"`
}

type errorMessage struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, e error, code int) {
	res := errorResponse{}
	switch err := e.(type) {
	case *RequestValidationError:
		messages := make([]errorMessage, 0)
		for _, r := range err.Reasons {
			m := errorMessage{
				Field:   r.Field(),
				Message: r.Error(),
			}
			messages = append(messages, m)
		}
		res.Errors = messages
	case *JsonFormattingError:
		message := make([]errorMessage, 1)
		message[0] = errorMessage{
			Message: err.Error(),
		}
		res.Errors = message
	default:
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
