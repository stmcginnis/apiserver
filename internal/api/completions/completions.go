package completions

import (
	"fmt"
	"net/http"

	"github.com/stmcginnis/apiserver/internal/api/route"
)

func init() {
	route.RegisterEndpoint(&route.Endpoint{
		Method:  "GET",
		URI:     "/v1/chat/completions",
		Handler: chatCompletionHandler,
	})
}

func chatCompletionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "game")
}
