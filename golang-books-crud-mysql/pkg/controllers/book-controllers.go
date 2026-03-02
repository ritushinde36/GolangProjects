package controllers

import (
	"books-crud-mysql/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetBooksSQL()
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, _ := models.GetBookByIDSQL(id)
	json.NewEncoder(w).Encode(book)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book *models.Book
	json.NewDecoder(r.Body).Decode(&book)
	b := book.CreateBooksSQL()
	json.NewEncoder(w).Encode(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newBook *models.Book
	json.NewDecoder(r.Body).Decode(&newBook)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, db := models.GetBookByIDSQL(id)
	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}
	db.Save(&book)
	json.NewEncoder(w).Encode(book)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book := models.DeleteBookSQL(id)
	json.NewEncoder(w).Encode(book)

}
