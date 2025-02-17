package router

import (
	"net/http"

	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/handler"
)

func initProtectedRoutes(mux *http.ServeMux) {
	secret := []byte(config.Envs.JwtToken)

	mux.Handle("/api/v1/users/profile", Authenticate(secret)(http.HandlerFunc(handler.GetUserProfileByUsername)))
	mux.Handle("/api/v1/users/update", Authenticate(secret)(http.HandlerFunc(handler.UpdateUser)))
	mux.Handle("/api/v1/users/delete", Authenticate(secret)(http.HandlerFunc(handler.DeleteUser)))
	mux.Handle("/api/v1/titles/search", Authenticate(secret)(http.HandlerFunc(handler.GetTitlesSearch)))
	mux.Handle("/api/v1/comment", Authenticate(secret)(http.HandlerFunc(handler.AddComment)))
	mux.Handle("/api/v1/watchlater", Authenticate(secret)(http.HandlerFunc(handler.AddWatchLater)))
	mux.Handle("/api/v1/watched", Authenticate(secret)(http.HandlerFunc(handler.AddWatched)))
	mux.Handle("/api/v1/rate", Authenticate(secret)(http.HandlerFunc(rateHandler)))
	mux.Handle("/api/v1/favorite", Authenticate(secret)(http.HandlerFunc(favoriteHandler)))

}

func rateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.AddRating(w, r)
	case http.MethodDelete:
		handler.RemoveRating(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func favoriteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.AddFavorite(w, r)
	case http.MethodDelete:
		handler.RemoveFavorite(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
