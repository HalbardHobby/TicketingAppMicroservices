package routes

import (
	"fmt"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there")
}
