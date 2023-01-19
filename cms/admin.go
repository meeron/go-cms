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
const DefaultAdminUrl = "admin"

//go:embed admin
var adminEmbed embed.FS

func adminIndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(adminEmbed, fmt.Sprintf("%s/index.html", AdminAppDir))
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "text/html")
	tpl.Execute(w, struct{}{})
}

func configureAdmin(router *mux.Router) {
	r := router.PathPrefix("/" + DefaultAdminUrl).Subrouter()

	adminRootFs, err := fs.Sub(adminEmbed, AdminAppDir)
	if err != nil {
		log.Fatal(err)
	}
	adminStaticHandler := http.FileServer(http.FS(adminRootFs))

	r.HandleFunc("/", adminIndexHandler).Methods(http.MethodGet)
	r.HandleFunc("", adminIndexHandler).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.StripPrefix(fmt.Sprintf("/%s/", DefaultAdminUrl), adminStaticHandler))
}
