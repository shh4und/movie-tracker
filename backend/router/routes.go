package router

import (
	"net/http"

	"github.com/shh4und/movie-tracker/handler"
	"github.com/shh4und/movie-tracker/services"
)

func initRoutes(mux *http.ServeMux) {
	handler.InitHandler()
	services.InitServices()

	// Public routes
	initPublicRoutes(mux)

	// Protected routes
	initProtectedRoutes(mux)
}
