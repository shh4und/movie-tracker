package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// AddRating adds a rating for a movie
func AddRating(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	var ratingRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
		Rating      int    `json:"rating"`
	}
	if err := json.NewDecoder(r.Body).Decode(&ratingRequest); err != nil {
		logger.Errorf("error binding rating request: %v", err)
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
	err = tx.QueryRow(r.Context(), "SELECT id FROM titles WHERE imdb_id = $1", ratingRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to rate title: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Insert the rating into the user_ratings table
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		logger.Errorf("Failed to convert userID to int: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tx.Exec(r.Context(), `
        INSERT INTO user_ratings (user_id, title_id, rating)
        VALUES ($1, $2, $3)
    `, userIDInt, titleID, ratingRequest.Rating)
	if err != nil {
		logger.Errorf("Failed to insert rating: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "rating-added", ratingRequest)
}

func RemoveRating(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	var ratingRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
		Rating      int    `json:"rating"`
	}
	if err := json.NewDecoder(r.Body).Decode(&ratingRequest); err != nil {
		logger.Errorf("error binding rating request: %v", err)
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
	err = tx.QueryRow(r.Context(), "SELECT id FROM titles WHERE imdb_id = $1", ratingRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Remove the rating from the user_ratings table
	_, err = tx.Exec(r.Context(), "DELETE FROM user_ratings WHERE user_id = $1 AND title_id = $2", userID, titleID)
	if err != nil {
		logger.Errorf("Failed to remove rating: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "rating-removed", ratingRequest)
}
