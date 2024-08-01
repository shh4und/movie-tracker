package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/shh4und/movie-tracker/auth"
)

func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}
	query := "INSERT INTO users (username, email, password, minor) VALUES (@Username, @Email, @Password, @Minor)"
	args := pgx.NamedArgs{
		"Username": request.Username,
		"Email":    request.Email,
		"Password": hashedPassword,
		"Minor":    request.Minor,
	}

	_, err = dbpg.DB.Exec(ctx, query, args)
	if err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", request.Username)

}
