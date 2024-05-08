package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"CommandGet",
		strings.ToUpper("Get"),
		"/command",
		CommandGet,
	},

	Route{
		"CommandIdGet",
		strings.ToUpper("Get"),
		"/command/{id}",
		CommandIdGet,
	},

	Route{
		"CommandPost",
		strings.ToUpper("Post"),
		"/command",
		CommandPost,
	},

	Route{
		"StopIdPost",
		strings.ToUpper("Post"),
		"/stop/{id}",
		StopIdPost,
	},
}
