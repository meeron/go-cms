package cms

import (
	"fmt"
	"net/http"
)

type adminHandler struct{}

func (h adminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from admin")
}

func createAdminHandler() http.Handler {
	return adminHandler{}
}
