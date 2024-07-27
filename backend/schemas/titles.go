package schemas

import (
	"gorm.io/gorm"
)

// Title representa a tabela titles no banco de dados
type Title struct {
	gorm.Model
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	RatingsID  uint
	Metascore  string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

// Rating representa a tabela ratings no banco de dados
type Rating struct {
	gorm.Model
	TitleID uint
	Source  string
	Value   string
	Title   Title `gorm:"foreignKey:TitleID;constraint:OnDelete:CASCADE"`
}

// Search representa uma pesquisa com resultados
type Search struct {
	Titles       []Title
	TotalResults string
	Response     string
}
