package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFavorite(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(string)
	var favoriteRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
	}
	if err := ctx.ShouldBindJSON(&favoriteRequest); err != nil {
		logger.Errorf("error binding favorite request: %v", err)
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
	err = tx.QueryRow(ctx, "SELECT id FROM titles WHERE imdb_id = $1", favoriteRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	// Insert the favorite into the user_favorites table
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		logger.Errorf("Failed to convert userID to int: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tx.Exec(ctx, `
        INSERT INTO user_favorites (user_id, title_id)
        VALUES ($1, $2)
    `, userIDInt, titleID)
	if err != nil {
		logger.Errorf("Failed to insert favorite: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "favorite-added", favoriteRequest)
}

func RemoveFavorite(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(string)
	var favoriteRequest struct {
		TitleIMDbID string `json:"title_imdbID"`
	}
	if err := ctx.ShouldBindJSON(&favoriteRequest); err != nil {
		logger.Errorf("error binding favorite request: %v", err)
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
	err = tx.QueryRow(ctx, "SELECT id FROM titles WHERE imdb_id = $1", favoriteRequest.TitleIMDbID).Scan(&titleID)
	if err != nil {
		// Title does not exist, insert it
		logger.Errorf("Failed to find title: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return

	}
	// Remove the favorite from the user_favorites table
	_, err = tx.Exec(ctx, "DELETE FROM user_favorites WHERE user_id = $1 AND title_id = $2", userID, titleID)
	if err != nil {
		logger.Errorf("Failed to remove favorite: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "favorite-removed", favoriteRequest)
}
