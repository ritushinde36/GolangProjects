package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ritushinde36/GolangProjects/golang-books-crud-api/books"
	"github.com/ritushinde36/GolangProjects/golang-books-crud-api/operations"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	books.AddBooks(10)
	router.HandleFunc("/books", operations.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", operations.GetBook).Methods("GET")
	router.HandleFunc("/books", operations.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", operations.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", operations.DeleteBook).Methods("DELETE")

	fmt.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

}
