package routes

import (
	"encoding/json"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		MaxAge: -1,
		Value:  "",
		Path:   "/",
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{})
}
