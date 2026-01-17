package operations

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ritushinde36/GolangProjects/golang-books-crud-api/books"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books.Books {
		if strconv.Itoa(book.BookId) == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook books.Book
	json.NewDecoder(r.Body).Decode(&newBook)
	books.Books = append(books.Books, newBook)
	json.NewEncoder(w).Encode(newBook)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books.Books {
		if strconv.Itoa(book.BookId) == params["id"] {
			books.Books = append(books.Books[:index], books.Books[index+1:]...)
			var newBook books.Book
			json.NewDecoder(r.Body).Decode(&newBook)
			books.Books = append(books.Books, newBook)
			json.NewEncoder(w).Encode(newBook)
			return
		}
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books.Books {
		if strconv.Itoa(book.BookId) == params["id"] {
			books.Books = append(books.Books[:index], books.Books[index+1:]...)
			json.NewEncoder(w).Encode(books.Books)
			return
		}
	}

}
