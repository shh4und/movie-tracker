package schemas

import "time"

// User represents the users table
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:50;unique;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:100;not null"`
	Minor     bool   `gorm:"not null"`
	FirstName string `gorm:"size:50"`
	LastName  string `gorm:"size:50"`
	PhotoURL  string `gorm:"size:255"`
	Status    string `gorm:"size:255"`
}

// UserRating represents the user_ratings table
type UserRating struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	MovieID uint `gorm:"not null"`
	Rating  int  `gorm:"check:rating >= 1 AND rating <= 10"`
	User    User `gorm:"constraint:OnDelete:CASCADE;"`
}

// UserFavorite represents the user_favorites table
type UserFavorite struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	MovieID uint `gorm:"not null"`
	User    User `gorm:"constraint:OnDelete:CASCADE;"`
}

// WatchLater represents the watch_later table
type WatchLater struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	MovieID uint `gorm:"not null"`
	User    User `gorm:"constraint:OnDelete:CASCADE;"`
}

// WatchedMovie represents the watched_movies table
type WatchedMovie struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	MovieID   uint `gorm:"not null"`
	WatchedOn time.Time
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"not null"`
	MovieID uint   `gorm:"not null"`
	Text    string `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Title   Title  `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE"`
}
