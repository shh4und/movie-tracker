package handler

import (
	"fmt"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
		return
	}

	tx, err := dbpg.DB.Begin(r.Context())

	if err != nil {
		logger.Errorf("error starting transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Rollback(r.Context())

	query := "DELETE FROM tracker.users WHERE id=$1"

	_, err = tx.Exec(r.Context(), query, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id:%s", id))
		return
	}

	if err := tx.Commit(r.Context()); err != nil {
		logger.Errorf("error committing transaction: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, "delete-user", id)

}
