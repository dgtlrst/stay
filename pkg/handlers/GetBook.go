package handlers

import (
	"crud/pkg/mocks"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// this handler obtains book information by supplied id
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	// find book from id
	for _, b := range mocks.Books {
		if b.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(b)
			break
		}
	}
}
