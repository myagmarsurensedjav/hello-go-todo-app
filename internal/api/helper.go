package api

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

func respondWithError(w http.ResponseWriter, err Error) {
	errorMessage := map[string]string{"error": err.Error}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(errorMessage)
	return
}
