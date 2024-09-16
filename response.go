package http_utils

import (
	"encoding/json"
	"net/http"
)

func WriteOkJson[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteBadRequestWithJson[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ValidateRequestQuery(r *http.Request, key string) bool {
	return r.URL.Query().Has(key)
}
