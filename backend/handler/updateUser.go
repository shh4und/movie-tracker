package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/schemas"
)

func UpdateUser(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	request := UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := schemas.User{}
	updateFields := schemas.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		PhotoURL:  request.PhotoURL,
		Status:    request.Status,
		LastName:  request.LastName,
	}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found on the database", id))
		return
	}
	if err := db.Model(&user).Updates(&updateFields).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error at update operation, user id: %s", id))
		return
	}

	sendSuccess(ctx, "update-user", user)

}
