package handlers

import (
	"crud/pkg/mocks"
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}
