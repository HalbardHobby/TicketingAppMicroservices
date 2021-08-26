package routes

import (
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there")
}
