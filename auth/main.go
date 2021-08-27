package main

import (
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/routes"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	ne := &errors.NotFoundError{Reason: "Resource not found", Code: http.StatusNotFound}
	errors.JsonError(w, ne)
}

func main() {
	data.ConnectDB()

	r := mux.NewRouter()
	s := r.PathPrefix("/api/users").Subrouter()

	s.HandleFunc("/currentuser", routes.CurrentUser).Methods("GET")
	s.HandleFunc("/signup", routes.SignUp).Methods("POST")
	s.HandleFunc("/signin", routes.SignIn).Methods("POST")
	s.HandleFunc("/signout", routes.SignOut).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notFound)

	log.Print("Listening on port 4200!")
	log.Fatal(http.ListenAndServe(":4200", r))
}
