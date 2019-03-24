package router

import (
	"net/http"

	handler "github.com/port-project/api-service/port/delivery/rest"
)

// NewRouter is used to create router instance
func NewRouter() *router {

	handler := handler.NewHandler()
	router := &router{handler}
	return router
}

// Router is used to handle events from http server
type router struct {
	http.Handler
}
