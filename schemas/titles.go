package structs

import (
	"gorm.io/gorm"
)

type Title struct {
	gorm.Model
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
}
