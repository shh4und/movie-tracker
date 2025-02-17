package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/shh4und/movie-tracker/auth"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	var request UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Errorf("request binding error: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	var updateFields []string
	var updateValues []interface{}
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if request.Username != "" {
		updateFields = append(updateFields, "username=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.Username)

	}
	if request.Email != "" {
		updateFields = append(updateFields, "email=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.Email)

	}
	if request.Password != "" {
		updateFields = append(updateFields, "password=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, hashedPassword)

	}
	if request.FirstName != "" {
		updateFields = append(updateFields, "first_name=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.FirstName)

	}
	if request.LastName != "" {
		updateFields = append(updateFields, "last_name=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.LastName)

	}
	if request.PhotoURL != "" {
		updateFields = append(updateFields, "photo_url=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.PhotoURL)

	}
	if request.Status != "" {
		updateFields = append(updateFields, "status=$"+strconv.Itoa(len(updateValues)+1))
		updateValues = append(updateValues, request.Status)

	}
	tx, err := dbpg.DB.Begin(r.Context())

	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	query := fmt.Sprintf("UPDATE tracker.users SET %s WHERE id=$%d RETURNING *", strings.Join(updateFields, ", "), len(updateValues)+1)

	updateValues = append(updateValues, id)

	_, err = tx.Query(r.Context(), query, updateValues...)
	if err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "update-user", request.Username)

}
