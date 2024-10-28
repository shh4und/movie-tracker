package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
)

func UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	request := UpdateUserRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("request binding error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var updateFields []string
	var updateValues []interface{}
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
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

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d RETURNING *", strings.Join(updateFields, ", "), len(updateValues)+1)

	updateValues = append(updateValues, id)

	_, err = dbpg.DB.Query(ctx, query, updateValues...)
	if err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "update-user", request.Username)

}
