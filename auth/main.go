package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/errors"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/middleware"
	"github.com/HalbardHobby/TicketingAppMicroservices/auth/routes"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	errors.NotFoundError(w, "Resource not Found")
}

func main() {
	_, present := os.LookupEnv("JWT_KEY")
	if !present {
		log.Panic("JWT_KEY environment variable not set")
	}

	data.ConnectDB()

	r := mux.NewRouter()
	s := r.PathPrefix("/api/users").Subrouter()
	s1 := s.PathPrefix("/").Subrouter()
	s2 := s.PathPrefix("/").Subrouter()

	s1.HandleFunc("/signup", routes.SignUp).Methods("POST")
	s1.HandleFunc("/signin", routes.SignIn).Methods("POST")
	s2.HandleFunc("/currentuser", routes.CurrentUser).Methods("GET")
	s.HandleFunc("/signout", routes.SignOut).Methods("POST")

	s1.Use(middleware.ValidateUserInput)
	s2.Use(middleware.GetCurrentUser)

	r.NotFoundHandler = http.HandlerFunc(notFound)

	log.Print("Listening on port 4200!")
	log.Fatal(http.ListenAndServe(":4200", r))
}
