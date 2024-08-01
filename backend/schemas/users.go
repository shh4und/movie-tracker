package schemas

import "time"

// User represents the users table
type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	Minor     bool
	FirstName string
	LastName  string
	PhotoURL  string
	Status    string
}

// UserRating represents the user_ratings table
type UserRating struct {
	ID      uint
	UserID  uint
	MovieID uint
	Rating  int
}

// UserFavorite represents the user_favorites table
type UserFavorite struct {
	ID      uint
	UserID  uint
	MovieID uint
}

// WatchLater represents the watch_later table
type WatchLater struct {
	ID      uint
	UserID  uint
	MovieID uint
}

// WatchedMovie represents the watched_movies table
type WatchedMovie struct {
	ID        uint
	UserID    uint
	MovieID   uint
	WatchedOn time.Time
}

type Comment struct {
	ID      uint
	UserID  uint
	MovieID uint
	Text    string
}
