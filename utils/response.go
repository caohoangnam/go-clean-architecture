package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": msg,
	}
	json.NewEncoder(w).Encode(body)
}
