package rest

import "net/http"

// Handler is used to parse incoming data
type Handler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
