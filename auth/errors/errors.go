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

type NotFoundError struct {
	Reason string
	Code   int
}

func (e *NotFoundError) Error() string {
	return e.Reason
}

func (e *NotFoundError) StatusCode() int {
	return e.Code
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

type JsonFormattingError struct {
	Reason string
	Code   int
}

func (e *JsonFormattingError) Error() string {
	return e.Reason
}

func (e *JsonFormattingError) StatusCode() int {
	return e.Code
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

type BadRequestError struct {
	Reason string
	Code   int
}

func (e *BadRequestError) Serialize() ErrorResponse {
	res := ErrorResponse{}
	message := make([]ErrorMessage, 1)
	message[0] = ErrorMessage{
		Message: e.Error(),
	}
	res.Errors = message
	return res
}

func (e *BadRequestError) Error() string {
	return e.Reason
}

func (e *BadRequestError) StatusCode() int {
	return e.Code
}

type ErrorResponse struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func JsonError(w http.ResponseWriter, e SerializableError) {
	res := e.Serialize()
	log.Println(e.Error())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(e.StatusCode())
	json.NewEncoder(w).Encode(res)
}
