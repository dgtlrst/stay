package handlers

import (
	"crud/pkg/mocks"
	"crud/pkg/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// parse user input dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	// iterate over all mock books
	for index, book := range mocks.Books {
		if book.Id == id {
			// update book
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			// send response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
