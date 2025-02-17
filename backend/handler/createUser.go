package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shh4und/movie-tracker/auth"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Errorf("error binding request: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate request
	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Begin transaction
	tx, err := dbpg.DB.Begin(r.Context())
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	// Insert user into database
	query := "INSERT INTO tracker.users (username, email, password) VALUES ($1, $2, $3)"
	_, err = tx.Exec(r.Context(), query, request.Username, request.Email, hashedPassword)
	if err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Commit transaction
	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "create-user", request.Username)
}
