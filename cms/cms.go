package cms

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var logger = log.Default()

type CmsApp struct {
	router *mux.Router
}

func (app *CmsApp) Run(addr string) error {
	logger.Printf("Listening on %s ...\n", addr)

	http.Handle("/", app.router)

	return http.ListenAndServe(addr, nil)
}

func New() CmsApp {
	router := mux.NewRouter()

	configureAdmin(router)
	configureNodes(router)

	return CmsApp{
		router,
	}
}
