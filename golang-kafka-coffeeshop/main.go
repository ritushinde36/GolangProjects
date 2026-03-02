package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritushinde36/GolangProjects/golang-kafka-coffeeshop/order_operations"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", order_operations.CreateOrder).Methods("POST")

	fmt.Println("starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}

	//consumer

}
