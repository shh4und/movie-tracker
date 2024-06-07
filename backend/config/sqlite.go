package config

import (
	"os"

	schemas "github.com/shh4und/movie-tracker/schemas"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/move_tracker.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Try to create the database dir and file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// Create DB and Connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the struct (schemas)
	err = db.AutoMigrate(&schemas.Title{}, &schemas.Rating{}, &schemas.User{}, &schemas.UserFavorite{}, &schemas.UserRating{}, &schemas.WatchLater{}, &schemas.WatchedMovie{})
	if err != nil {
		logger.Errorf("auto migrate error: %v", err)
		return nil, err
	}

	return db, nil
}
