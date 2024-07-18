package files

import (
	"net/http"

	"github.com/stmcginnis/apiserver/internal/api/route"
)

func init() {
	route.RegisterEndpoint(&route.Endpoint{
		Method:  "GET",
		URI:     "/v1/files",
		Handler: listFilesHandler,
	})
	route.RegisterEndpoint(&route.Endpoint{
		Method:  "POST",
		URI:     "/v1/files",
		Handler: uploadFilesHandler,
	})
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "unsupported", http.StatusNotFound)
}

func uploadFilesHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "unsupported", http.StatusNotFound)
}
