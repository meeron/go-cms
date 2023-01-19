package cms

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const AdminAppDir = "admin"

//go:embed admin
var adminEmbed embed.FS

func adminIndexHandler(config *CmsConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFS(adminEmbed, fmt.Sprintf("%s/index.html", AdminAppDir))
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "text/html")
		tpl.Execute(w, struct{ AdminRootUrl string }{
			AdminRootUrl: config.AdminRootUrl,
		})
	}
}

func configureAdmin(router *mux.Router, config *CmsConfig) {
	r := router.PathPrefix("/" + config.AdminRootUrl).Subrouter()

	adminRootFs, err := fs.Sub(adminEmbed, AdminAppDir)
	if err != nil {
		log.Fatal(err)
	}
	adminStaticHandler := http.FileServer(http.FS(adminRootFs))

	r.HandleFunc("/", adminIndexHandler(config)).Methods(http.MethodGet)
	r.HandleFunc("", adminIndexHandler(config)).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.StripPrefix(fmt.Sprintf("/%s/", config.AdminRootUrl), adminStaticHandler))
}
