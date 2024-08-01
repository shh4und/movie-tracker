package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	query := "DELETE FROM users WHERE id=@ID"
	args := pgx.NamedArgs{
		"ID": id,
	}

	// var user schemas.User

	_, err := dbpg.DB.Exec(ctx, query, args)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id:%s", id))
		return
	}

	sendSuccess(ctx, "delete-user", id)

}
