package route

import (
	"fmt"
	"net/http"
)

var registeredEndpoints []*Endpoint

type Endpoint struct {
	Method  string
	URI     string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s %s", endpoint.Method, endpoint.URI)
}

func RegisterEndpoint(endpoint *Endpoint) {
	registeredEndpoints = append(registeredEndpoints, endpoint)
}

func Endpoints() []*Endpoint {
	return registeredEndpoints
}
