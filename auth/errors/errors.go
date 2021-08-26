package errors

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SerializableError interface {
	error
	Serialize() ErrorResponse
}

type NotFoundError struct {
	Reason string
}

func (e *NotFoundError) Error() string {
	return e.Reason
}

func (e *NotFoundError) Serialize() ErrorResponse {
	res := ErrorResponse{}
	message := make([]ErrorMessage, 1)
	message[0] = ErrorMessage{
		Message: e.Error(),
	}
	res.Errors = message
	return res
}

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

func (e *RequestValidationError) Serialize() ErrorResponse {
	res := ErrorResponse{}
	messages := make([]ErrorMessage, 0)
	for _, r := range e.Reasons {
		m := ErrorMessage{
			Field:   r.Field(),
			Message: r.Error(),
		}
		messages = append(messages, m)
	}
	res.Errors = messages
	return res
}

type JsonFormattingError struct {
	Reason string
}

func (e *JsonFormattingError) Serialize() ErrorResponse {
	res := ErrorResponse{}
	message := make([]ErrorMessage, 1)
	message[0] = ErrorMessage{
		Message: e.Error(),
	}
	res.Errors = message
	return res
}

func (e *JsonFormattingError) Error() string {
	return e.Reason
}

type ErrorResponse struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, e SerializableError, code int) {
	res := e.Serialize()
	log.Println(e.Error())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
