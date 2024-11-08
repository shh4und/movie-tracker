package models

import "time"

// Title represents the titles table in the database and the API response
type Title struct {
	ID         uint      `json:"id"`
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	IMDBRating string    `json:"imdbRating"`
	IMDBID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	Production string    `json:"Production"`
	Response   string    `json:"Response"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Search represents a search with results the API response
type Search struct {
	Titles       []Title `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

// User represents the users table
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoURL  string `json:"photo_url"`
	Status    string `json:"status"`
}

// UserRating represents the user_ratings table
type UserRating struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	TitleID uint `json:"title_id"`
	Rating  int  `json:"rating"`
}

// UserFavorite represents the user_favorites table
type UserFavorite struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	TitleID uint `json:"title_id"`
}

// WatchLater represents the watch_later table
type WatchLater struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	TitleID uint `json:"title_id"`
}

// WatchedMovie represents the watched_movies table
type WatchedMovie struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	TitleID   uint      `json:"title_id"`
	WatchedOn time.Time `json:"watched_on"`
}

// Comment represents the comments table
type Comment struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	TitleID uint   `json:"title_id"`
	Text    string `json:"text"`
}
