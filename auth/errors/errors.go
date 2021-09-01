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
	StatusCode() int
	Serialize() ErrorResponse
}

type RequestValidationError struct {
	Reasons []validator.FieldError
	Code    int
}

func (e *RequestValidationError) Error() string {
	var err strings.Builder
	for _, e := range e.Reasons {
		err.WriteString(e.Error() + "\n")
	}
	return err.String()
}

func (e *RequestValidationError) StatusCode() int {
	return e.Code
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

type RequestError struct {
	Reason string
	Code   int
}

func (e *RequestError) Serialize() ErrorResponse {
	res := ErrorResponse{}
	message := make([]ErrorMessage, 1)
	message[0] = ErrorMessage{
		Message: e.Error(),
	}
	res.Errors = message
	return res
}

func (e *RequestError) Error() string {
	return e.Reason
}

func (e *RequestError) StatusCode() int {
	return e.Code
}

type ErrorResponse struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func BadRequestError(rw http.ResponseWriter, reason string) {
	re := &RequestError{
		Reason: reason,
		Code:   400,
	}
	JsonError(rw, re)
}

func NotAuthorizedError(rw http.ResponseWriter, reason string) {
	re := &RequestError{
		Reason: reason,
		Code:   401,
	}
	JsonError(rw, re)
}

func NotFoundError(rw http.ResponseWriter, reason string) {
	re := &RequestError{
		Reason: reason,
		Code:   404,
	}
	JsonError(rw, re)
}

func JsonFormattingError(rw http.ResponseWriter, reason string) {
	re := &RequestError{
		Reason: reason,
		Code:   400,
	}
	JsonError(rw, re)
}

func JsonError(w http.ResponseWriter, e SerializableError) {
	res := e.Serialize()
	log.Println(e.Error())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(e.StatusCode())
	json.NewEncoder(w).Encode(res)
}
