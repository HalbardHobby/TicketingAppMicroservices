package routes

import (
	"encoding/json"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/middleware"
)

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(middleware.UserContextKey)
	if currentUser == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"currentUser": "null"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(currentUser)
}
