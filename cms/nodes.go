package cms

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var sampleNodes = [3]string{"", "node", "child"}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	statusCode := 404

	defer func() {
		elapsed := time.Since(start)
		logger.Printf("%v %v %v %v", r.Method, r.RequestURI, statusCode, elapsed)
	}()

	nodes := strings.Split(r.RequestURI, "/")

	// TODO: Find and render node
	if len(nodes) == len(sampleNodes) && nodes[1] == sampleNodes[1] && nodes[2] == sampleNodes[2] {
		statusCode = 200

		fmt.Fprintf(w, "Hello default node handler. Path: %s", r.RequestURI)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Page not found")
}

func configureNodes(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(defaultHandler)
}
