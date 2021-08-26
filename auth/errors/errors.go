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

func JsonError(w http.ResponseWriter, e error, code int) {
	formattedErr := make(map[string]interface{})

	switch err := e.(type) {
	case *RequestValidationError:
		formErrs := make([]interface{}, 0)
		for _, r := range err.Reasons {
			m := make(map[string]string)
			m["message"] = r.Error()
			m["field"] = r.Field()
			formErrs = append(formErrs, m)
		}
		formattedErr["errors"] = formErrs
	case *JsonFormattingError:
		formErrs := make([]interface{}, 1)
		m := make(map[string]string)
		m["message"] = e.Error()
		formErrs[0] = m
		formattedErr["errors"] = formErrs
	default:
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(formattedErr)
}
