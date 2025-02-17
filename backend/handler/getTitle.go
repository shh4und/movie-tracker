package handler

import (
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/shh4und/movie-tracker/models"
	"github.com/shh4und/movie-tracker/services"
)

// need fixes
// handler for fetching a list of a title search
func GetTitle(w http.ResponseWriter, r *http.Request) {
	titleName := r.URL.Query().Get("title")

	normalizedTitleName := strings.ToLower(titleName) // Normalizar o título para minúsculas

	tx, err := dbpg.DB.Begin(r.Context())
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	// Check if the title exists in the titles table
	var title *models.Title
	rows, err := tx.Query(r.Context(), "SELECT * FROM titles WHERE LOWER(title) = $1", normalizedTitleName)
	if err != nil {
		logger.Errorf("error querying titles: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	title, err = pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[models.Title])
	if err != nil {
		title, err = services.FetchTitleFromOMDB(normalizedTitleName)
		if err != nil {
			logger.Errorf("error at FetchTitleFromOMDB: %v", err.Error())
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Insert the title into the titles table
		err = tx.QueryRow(r.Context(), `
            INSERT INTO titles (title, year, rated, released, runtime, genre, director, writer, actors, plot, language, country, awards, poster, imdb_rating, imdb_id, type, production, response, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
            RETURNING id
        `, title.Title, title.Year, title.Rated, title.Released, title.Runtime, title.Genre, title.Director, title.Writer,
			title.Actors, title.Plot, title.Language, title.Country, title.Awards, title.Poster, title.IMDBRating, title.IMDBID,
			title.Type, title.Production, title.Response).Scan(&title.ID)
		if err != nil {
			logger.Errorf("Failed to insert title: %v", err.Error())
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		logger.Infof("the title: %v have been cached", normalizedTitleName)
	}

	if err = tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(w, "get-title", title)
}
