package handler

import (
	"fmt"
)

func errParamIsRequired(name, tp string) error {
	return fmt.Errorf("( param: %s | type: %s ) is required", name, tp)
}

// Create User
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Minor    bool   `json:"minor"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Username == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	if len(r.Password) < 8 {
		return fmt.Errorf("the password must have at least 8 characters")
	}

	if r.Minor {
		return fmt.Errorf("the user must have at least 16 years of age")
	}

	return nil
}

// Update User Data
type UpdateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	PhotoURL  string `json:"photourl"`
	Status    string `json:"status"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Username == "" || r.Email == "" || r.Password == "" || r.FirstName == "" || r.LastName == "" || r.PhotoURL == "" || r.Status == "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}

type TitleRequest struct {
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

type SearchRequest struct {
	Titles       []TitleRequest `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
}

func (r *SearchRequest) Validate() error {
	if r.Response == "False" {
		return fmt.Errorf("movie not found")
	}
	return nil
}
