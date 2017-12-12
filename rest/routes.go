package rest

import (
	"net/http"
)

var routes = ConfiguredRoutes{
	Route{
		Name:        "GetUpdates",
		Method:      "GET",
		Pattern:     "/updates",
		HandlerFunc: HandleGetUpdates},
}

type ConfiguredRoutes []Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AuthHandler
}

type AuthHandler func(http.ResponseWriter, *http.Request)

func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// TODO: implement authentication

	fn(w, r)
}