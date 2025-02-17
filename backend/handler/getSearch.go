package handler

import (
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/shh4und/movie-tracker/models"
	"github.com/shh4und/movie-tracker/services"
)

// handler for fetching a list of a title search
func GetTitlesSearch(w http.ResponseWriter, r *http.Request) {
	titleName := r.URL.Query().Get("title")
	normalizedTitleName := strings.ToLower(titleName)

	tx, err := dbpg.DB.Begin(r.Context())
	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	// First try to find exact matches in database
	var titles []*models.Title
	rows, err := tx.Query(r.Context(), "SELECT * FROM tracker.titles WHERE LOWER(title) = $1", normalizedTitleName)
	if err != nil {
		logger.Errorf("error querying titles: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	// Collect all matching titles
	titles, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Title])
	if err != nil && err != pgx.ErrNoRows {
		logger.Errorf("error collecting rows: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If no titles found in DB, search OMDB
	if len(titles) == 0 {
		// Search OMDB for possible matches
		search, err := services.FetchSearchFromOMDB(normalizedTitleName)
		if err != nil {
			logger.Errorf("error searching OMDB: %v", err)
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Add limit to prevent overwhelming response
		if len(search.Titles) > 10 {
			search.Titles = search.Titles[:10]
		}
		// Store each found title in database
		for _, omdbTitle := range search.Titles {
			// Check if title already exists by IMDB ID
			var exists bool
			err = tx.QueryRow(r.Context(), "SELECT EXISTS(SELECT 1 FROM tracker.titles WHERE imdb_id = $1)", omdbTitle.IMDBID).Scan(&exists)
			if err != nil {
				logger.Errorf("error checking title existence: %v", err)
				continue
			}

			if !exists {
				// Fetch full details for this title
				fullTitle, err := services.FetchIMDbIDFromOMDB(omdbTitle.IMDBID)
				if err != nil {
					logger.Errorf("error fetching full title details: %v", err)
					continue
				}

				// Insert new title
				err = tx.QueryRow(r.Context(), `
                    INSERT INTO tracker.titles (
                        title, year, rated, released, runtime, genre, director, 
                        writer, actors, plot, language, country, awards, 
                        poster, imdb_rating, imdb_id, type, production, 
                        response, created_at, updated_at
                    ) VALUES (
                        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,
                        $13, $14, $15, $16, $17, $18, $19, 
                        CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
                    ) RETURNING id`,
					fullTitle.Title, fullTitle.Year, fullTitle.Rated,
					fullTitle.Released, fullTitle.Runtime, fullTitle.Genre,
					fullTitle.Director, fullTitle.Writer, fullTitle.Actors,
					fullTitle.Plot, fullTitle.Language, fullTitle.Country,
					fullTitle.Awards, fullTitle.Poster, fullTitle.IMDBRating,
					fullTitle.IMDBID, fullTitle.Type, fullTitle.Production,
					fullTitle.Response).Scan(&fullTitle.ID)

				if err != nil {
					logger.Errorf("error inserting title: %v", err)
					continue
				}
				titles = append(titles, fullTitle)
			}
		}
	}
	// Add limit to prevent overwhelming response
	if len(titles) > 10 {
		titles = titles[:10]
	}

	if err = tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "get-title", titles)
}
