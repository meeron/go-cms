package cms

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from admin. Path %v", r.RequestURI)
}

func configureAdmin(router *mux.Router) {
	r := router.PathPrefix("/admin").Subrouter()

	r.HandleFunc("/", adminHandler).Methods(http.MethodGet)
	r.HandleFunc("", adminHandler).Methods(http.MethodGet)
}
