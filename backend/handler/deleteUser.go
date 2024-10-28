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

	tx, err := dbpg.DB.Begin(ctx)

	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(ctx)

	query := "DELETE FROM users WHERE id=$1"

	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id:%s", id))
		return
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "delete-user", id)

}
