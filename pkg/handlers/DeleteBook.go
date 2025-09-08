package handlers

import (
	"crud/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// getid from user for book to delete
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// iterate over all mock books and find book with dynamic id
	for index, book := range mocks.Books {
		if book.Id == id {
			// remove book from mock
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			// send response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte("Deleted"))
			break
		}
	}
}
