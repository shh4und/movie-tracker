package router

import (
	"net/http"
)

func Init() {
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("../frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Serve main HTML files
	mux.Handle("/index", http.FileServer(http.Dir("../frontend/static/index.html")))
	mux.Handle("/register", http.FileServer(http.Dir("../frontend/static/register.html")))
	mux.Handle("/login", http.FileServer(http.Dir("../frontend/static/login.html")))

	// Initialize routes
	initRoutes(mux)

	// Run server
	http.ListenAndServe(":8080", mux)
}
