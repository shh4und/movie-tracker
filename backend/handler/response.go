package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   false,
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(w http.ResponseWriter, op string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("operation from handler %s: successfully done", op),
		"success": true,
		"data":    data,
	})
}
