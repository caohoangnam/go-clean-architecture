package utils

import (
	"net/http"
)

func ResponseSuccess(w http.ResponseWrite, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncode(w).Encode(body)
}

func ResponseError(w http.ResponseWrite, code int, msg string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": msg,
	}
	json.NewEncode(w).Encode(w)
}
