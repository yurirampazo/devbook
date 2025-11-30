package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Puts routes into router
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, feedPostRoutes...)

	for _, route := range routes {
		if route.RequireAuth {		
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),).Methods(route.Method)
		}

		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}
	
	return r
}
