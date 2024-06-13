package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/schemas"
)

func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := schemas.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Minor:    request.Minor,
	}

	if err := db.Create(&newUser).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, newUser)

}
