package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/models"
	"github.com/shh4und/movie-tracker/services"
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
	var title models.Title
	err = tx.QueryRow(ctx, "SELECT id FROM titles WHERE imdb_id = $1", ratingRequest.TitleIMDbID).Scan(&title.ID)
	if err != nil {
		// Title does not exist, insert it
		// Fetch title details from OMDB API
		title, err := services.FetchIMDbIDFromOMDB(ratingRequest.TitleIMDbID)
		if err != nil {
			logger.Errorf("Failed to fetch title details: %v", err)
			sendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		// Insert the title into the titles table
		err = tx.QueryRow(ctx, `
            INSERT INTO titles (title, year, rated, released, runtime, genre, director, writer, actors, plot, language, country, awards, poster, imdb_rating, imdb_id, type, production, response, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
            RETURNING id
        `, title.Title, title.Year, title.Rated, title.Released, title.Runtime, title.Genre, title.Director, title.Writer,
			title.Actors, title.Plot, title.Language, title.Country, title.Awards, title.Poster, title.IMDBRating, title.IMDBID,
			title.Type, title.Production, title.Response).Scan(&title.ID)
		if err != nil {
			logger.Errorf("Failed to insert title: %v", err.Error())
			sendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
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
    `, userIDInt, title.ID, ratingRequest.Rating)
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

	sendSuccess(ctx, "rating-added", ratingRequest.Rating)

}
