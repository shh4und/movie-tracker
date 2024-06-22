package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler for fetching a list of a title search
func GetSearch(ctx *gin.Context) {
	title := ctx.Query("title")
	apiKey := apiKEY
	if apiKey == "" {
		sendError(ctx, http.StatusInternalServerError, fmt.Errorf("apiKEY not set").Error())
		return
	}

	if title == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("title", "query-param").Error())
		return
	}
	omdbAPIBaseURL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s", apiKey, title) //move to config file later
	resp, err := http.Get(omdbAPIBaseURL)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		sendError(ctx, resp.StatusCode, resp.Status)
		return
	}

	request := SearchRequest{}

	if err := json.NewDecoder(resp.Body).Decode(&request); err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err = request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "search-titles", request)
}
