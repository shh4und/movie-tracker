package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/schemas"
)

func CreateUser(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
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
