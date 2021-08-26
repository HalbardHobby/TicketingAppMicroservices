package errors

import (
	"encoding/json"
	"net/http"
)

type RequestValidationError struct {
}

func (r *RequestValidationError) Error() string {
	return ""
}

func JsonError(w http.ResponseWriter, v interface{}, code int) {
	m := make(map[string]interface{})
	m["message"] = v

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(m)
}
