package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	var favoriteRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&favoriteRequest); err != nil {
		logger.Errorf("error binding favorite request: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := dbpg.DB.Begin(r.Context())
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	// Check if the title exists in the titles table
	var titleID uint
	err = tx.QueryRow(r.Context(), "SELECT id FROM titles WHERE imdb_id = $1", favoriteRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Insert the favorite into the user_favorites table
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		logger.Errorf("Failed to convert userID to int: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tx.Exec(r.Context(), `
        INSERT INTO user_favorites (user_id, title_id)
        VALUES ($1, $2)
    `, userIDInt, titleID)
	if err != nil {
		logger.Errorf("Failed to insert favorite: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "favorite-added", favoriteRequest)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	var favoriteRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&favoriteRequest); err != nil {
		logger.Errorf("error binding favorite request: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := dbpg.DB.Begin(r.Context())
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	// Check if the title exists in the titles table
	var titleID uint
	err = tx.QueryRow(r.Context(), "SELECT id FROM titles WHERE imdb_id = $1", favoriteRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Remove the favorite from the user_favorites table
	_, err = tx.Exec(r.Context(), "DELETE FROM user_favorites WHERE user_id = $1 AND title_id = $2", userID, titleID)
	if err != nil {
		logger.Errorf("Failed to remove favorite: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "favorite-removed", favoriteRequest)
}
