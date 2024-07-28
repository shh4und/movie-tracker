package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/schemas"
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
	query := "INSERT INTO users (username, email, password, minor) VALUES ($1, $2, $3, $4)"
	newUser := schemas.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		Minor:    request.Minor,
	}

	rows, err := dbpq.Query(ctx, query, request.Username, request.Email, hashedPassword, request.Minor)
	if err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		fmt.Printf("******* values: %v", values)
	}
	dbpq.Close()
	if err := db.Create(&newUser).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", newUser)

}
