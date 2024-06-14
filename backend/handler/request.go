package handler

import (
	"fmt"
)

func errParamIsRequired(name, tp string) error {
	return fmt.Errorf("( param: %s | type: %s )is required", name, tp)
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
