package routes

import (
	"fmt"
	"net/http"
)

func CurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there")
}
