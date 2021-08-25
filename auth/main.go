package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":4200", nil)
	fmt.Print("Listening on port 4200!")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
