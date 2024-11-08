package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/shh4und/movie-tracker/models"
	"github.com/shh4und/movie-tracker/services"
)

// handler for fetching a list of a title search
func GetTitle(ctx *gin.Context) {
	titleName := ctx.Query("title")

	normalizedTitleName := strings.ToLower(titleName) // Normalizar o título para minúsculas

	tx, err := dbpg.DB.Begin(ctx)
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(ctx)

	// Check if the title exists in the titles table
	var title *models.Title
	rows, err := tx.Query(ctx, "SELECT * FROM titles WHERE LOWER(title) = $1", normalizedTitleName)
	if err != nil {
		logger.Errorf("error querying titles: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	title, err = pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[models.Title])
	if err != nil {
		fmt.Println("\nERROR at CollectOneRow: \n" + err.Error())
		title, err = services.FetchTitleFromOMDB(normalizedTitleName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error at FetchTitleFromOMDB": err.Error()})
			return
		}
		// Insert the title into the titles table
		_, err = tx.Exec(ctx, `
            INSERT INTO titles (title, year, rated, released, runtime, genre, director, writer, actors, plot, language, country, awards, poster, imdb_rating, imdb_id, type, production, response, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
        `, strings.ToLower(title.Title), title.Year, title.Rated, title.Released, title.Runtime, title.Genre, title.Director, title.Writer,
			title.Actors, title.Plot, title.Language, title.Country, title.Awards, title.Poster, title.IMDBRating, title.IMDBID,
			title.Type, title.Production, title.Response)
		if err != nil {
			logger.Errorf("Failed to insert title: %v", err.Error())
			sendError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

	}

	if err = tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "get-title", title)
}
