package handler

import (
	"fmt"
	"net/http"

	"github.com/shh4und/movie-tracker/models"
)

func GetUserProfileByUsername(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")

	if username == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("username", "query-param").Error())
		return
	}

	query := "SELECT username FROM tracker.users WHERE username=$1"

	var user models.User

	err := dbpg.DB.QueryRow(r.Context(), query, username).Scan(&user.Username)
	if err != nil {
		sendError(w, http.StatusNotFound, fmt.Sprintf("user with username: %s not found on the database", username))
		return
	}

	sendSuccess(w, "get-user-username", user.Username)

}
