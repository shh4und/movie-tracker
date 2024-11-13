package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddRating adds a rating for a movie
func AddRating(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(string)
	var ratingRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
		Rating      int    `json:"rating"`
	}
	if err := ctx.ShouldBindJSON(&ratingRequest); err != nil {
		logger.Errorf("error binding rating request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := dbpg.DB.Begin(ctx)
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(ctx)

	// Check if the title exists in the titles table
	var titleID uint
	err = tx.QueryRow(ctx, "SELECT id FROM titles WHERE imdb_id = $1", ratingRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		// Fetch title details from OMDB API
		// Title does not exist, insert it
		logger.Errorf("Failed to rate title: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Insert the rating into the user_ratings table
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		logger.Errorf("Failed to convert userID to int: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tx.Exec(ctx, `
        INSERT INTO user_ratings (user_id, title_id, rating)
        VALUES ($1, $2, $3)
    `, userIDInt, titleID, ratingRequest.Rating)
	if err != nil {
		logger.Errorf("Failed to insert rating: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	if err := tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "rating-added", ratingRequest)

}

func RemoveRating(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(string)
	var ratingRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
		Rating      int    `json:"rating"`
	}
	if err := ctx.ShouldBindJSON(&ratingRequest); err != nil {
		logger.Errorf("error binding rating request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := dbpg.DB.Begin(ctx)
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(ctx)

	// Check if the title exists in the titles table
	var titleID uint
	err = tx.QueryRow(ctx, "SELECT id FROM titles WHERE imdb_id = $1", ratingRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return

	}
	// Remove the favorite from the user_favorites table
	_, err = tx.Exec(ctx, "DELETE FROM user_ratings WHERE user_id = $1 AND title_id = $2", userID, titleID)
	if err != nil {
		logger.Errorf("Failed to remove rating: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "rating-removed", ratingRequest)
}
