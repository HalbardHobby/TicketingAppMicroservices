package main

import (
	"log"
	"net/http"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/users").Subrouter()

	s.HandleFunc("/currentuser", routes.CurrentUser).Methods("GET")

	log.Print("Listening on port 4200!")
	log.Fatal(http.ListenAndServe(":4200", r))
}
