package routes

import (
	"fmt"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there")
}
