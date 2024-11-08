package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shh4und/movie-tracker/models"
)

// FetchTitleFromOMDB fetches title details from the OMDB API
func FetchTitleFromOMDB(titleName string) (*models.Title, error) {

	if apiKEY == "" {
		return nil, fmt.Errorf("API key is required")
	}

	omdbAPIBaseURL := "http://www.omdbapi.com/" + "?apikey=" + apiKEY + "&t=" + titleName
	resp, err := http.Get(omdbAPIBaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OMDB API request failed with status: %s", resp.Status)
	}

	var title models.Title
	if err := json.NewDecoder(resp.Body).Decode(&title); err != nil {
		return nil, err
	}

	return &title, nil
}

// FetchTitleFromOMDB fetches imdbid title details from the OMDB API
func FetchIMDbIDFromOMDB(titleName string) (*models.Title, error) {

	if apiKEY == "" {
		return nil, fmt.Errorf("API key is required")
	}

	omdbAPIBaseURL := "http://www.omdbapi.com/" + "?apikey=" + apiKEY + "&i=" + titleName
	resp, err := http.Get(omdbAPIBaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OMDB API request failed with status: %s", resp.Status)
	}

	var title models.Title
	if err := json.NewDecoder(resp.Body).Decode(&title); err != nil {
		return nil, err
	}

	return &title, nil
}

// FetchSearchFromOMDB fetches a search titles details from the OMDB API
func FetchSearchFromOMDB(titleName string) (*models.Search, error) {

	if apiKEY == "" {
		return nil, fmt.Errorf("API key is required")
	}

	omdbAPIBaseURL := "http://www.omdbapi.com/" + "?apikey=" + apiKEY + "&s=" + titleName
	resp, err := http.Get(omdbAPIBaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OMDB API request failed with status: %s", resp.Status)
	}

	var search models.Search
	if err := json.NewDecoder(resp.Body).Decode(&search); err != nil {
		return nil, err
	}

	return &search, nil
}
