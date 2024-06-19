package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/schemas"
)

func GetUserProfileByUsername(ctx *gin.Context) {
	username := ctx.Query("username")

	if username == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("username", "query-param").Error())
		return
	}
	user := schemas.User{}

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with username: %s not found on the database", username))
		return
	}

	sendSuccess(ctx, "get-user-username", user)

}
