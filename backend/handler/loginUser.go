package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/models"

	"github.com/gin-gonic/gin"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Errorf("request binding error: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	query := "SELECT id, username, password FROM tracker.users WHERE username=$1"

	err := dbpg.DB.QueryRow(r.Context(), query, request.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid username or password")
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(request.Password)) {
		sendError(w, http.StatusBadRequest, "invalid username or password")
		return
	}

	secret := []byte(config.Envs.JwtToken)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "login-user", gin.H{
		"token":    token,
		"userID":   user.ID,
		"username": user.Username,
	})

}
