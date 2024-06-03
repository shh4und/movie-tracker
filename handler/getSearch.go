package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/structs"
)

// handler for fetching a list of a title search
func GetSearch(ctx *gin.Context) {
	titleName := ctx.Param("title")
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "API_KEY not set"})
		return
	}
	omdbAPIBaseURL := "http://www.omdbapi.com/" + "?apikey=" + apiKey + "&s=" + titleName
	resp, err := http.Get(omdbAPIBaseURL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": resp.Status})
		return
	}

	var titles structs.Search
	if err := json.NewDecoder(resp.Body).Decode(&titles); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if titles.Response == "False" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	ctx.JSON(http.StatusOK, titles)
}
