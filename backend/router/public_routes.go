package router

import (
	"net/http"

	"github.com/shh4und/movie-tracker/handler"
)

func initPublicRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/register", handler.CreateUser)
	mux.HandleFunc("/api/v1/login", handler.LoginUser)
}
