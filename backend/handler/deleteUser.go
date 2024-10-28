package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	query := "DELETE FROM users WHERE id=$1"

	_, err := dbpg.DB.Exec(ctx, query, id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id:%s", id))
		return
	}

	sendSuccess(ctx, "delete-user", id)

}
