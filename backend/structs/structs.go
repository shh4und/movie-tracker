package structs

import "time"

type Title struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type Search struct {
	Titles       []Title `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

// User represents the users table
type User struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Minor    bool   `json:"minor"`
}

// UserRating represents the user_ratings table
type UserRating struct {
	UserID  uint `json:"userid"`
	MovieID uint `json:"movieid"`
	Rating  int  `json:"rate"`
}

// UserFavorite represents the user_favorites table
type UserFavorite struct {
	UserID  uint `json:"userid"`
	MovieID uint `json:"movieid"`
}

// WatchLater represents the watch_later table
type WatchLater struct {
	UserID  uint `json:"userid"`
	MovieID uint `json:"movieid"`
}

// WatchedMovie represents the watched_movies table
type WatchedMovie struct {
	UserID    uint      `json:"userid"`
	MovieID   uint      `json:"movieid"`
	WatchedOn time.Time `json:"watchedon"`
}

type UserProfile struct {
	UserID    uint   `json:"userid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	PhotoURL  string `json:"photourl"`
	Status    string `json:"status"`
}

type Comment struct {
	UserID  uint   `json:"userid"`
	MovieID uint   `json:"movieid"`
	Text    string `json:"comment"`
}
