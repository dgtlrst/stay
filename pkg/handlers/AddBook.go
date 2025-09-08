package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// append to book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// return status created
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
