package cms

import (
	"log"
	"net/http"
)

var logger = log.Default()

type CmsApp struct {
	server *http.ServeMux
}

func (app *CmsApp) Run(addr string) error {
	logger.Printf("Listening on %s ...\n", addr)
	return http.ListenAndServe(addr, app.server)
}

func New() CmsApp {
	server := http.NewServeMux()

	server.Handle("/admin", createAdminHandler())
	server.Handle("/", createNodeHandler())

	return CmsApp{
		server,
	}
}
