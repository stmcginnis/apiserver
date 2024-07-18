package api

import (
	_ "github.com/stmcginnis/apiserver/internal/api/completions"
	_ "github.com/stmcginnis/apiserver/internal/api/files"
	_ "github.com/stmcginnis/apiserver/internal/api/images"
	"github.com/stmcginnis/apiserver/internal/api/route"
)

func RegisteredEndpoints() []*route.Endpoint {
	return route.Endpoints()
}
