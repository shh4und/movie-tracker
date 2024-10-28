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

	query := "SELECT username FROM users WHERE username=$1"

	var user schemas.User

	err := dbpg.DB.QueryRow(ctx, query, username).Scan(&user.Username)
	if err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with username: %s not found on the database", username))
		return
	}

	sendSuccess(ctx, "get-user-username", user.Username)

}
