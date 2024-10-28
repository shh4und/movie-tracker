package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
)

func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("error binding request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	query := "INSERT INTO users (username, email, password, minor) VALUES ($1, $2, $3, $4)"

	_, err = dbpg.DB.Exec(ctx, query, request.Username, request.Email, hashedPassword, request.Minor)
	if err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", request.Username)

}
