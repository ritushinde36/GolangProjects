package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var items []string

func main() {
	fileServer := http.FileServer(http.Dir("./files"))
	http.Handle("/", fileServer)

	http.HandleFunc("/showitems", show_items)
	http.HandleFunc("/add", addItem)
	http.HandleFunc("/update", updateItem)
	http.HandleFunc("/delete", deleteItem)

	fmt.Printf("Starting localhost server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func show_items(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/showitems" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)

	}

	fmt.Fprintf(w, "\n\n-----------------YOUR TO DO LIST-----------------\n")
	for index, value := range items {
		var serialNumber int = 1
		serialNumber = serialNumber + index
		fmt.Fprintf(w, "%d . %v\n", serialNumber, value)
	}
	fmt.Fprintf(w, "-------------------------------------------------\n")

}

func addItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error : %v", err)
		return
	}
	item := r.FormValue("add_item")
	items = append(items, item)
	fmt.Fprintf(w, "Item : ' %v ' added successfully", item)

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error : %v", err)
		return
	}

	itemNumber := r.FormValue("update_item_number")
	updatedItem := r.FormValue("updated_item")

	ind, _ := strconv.Atoi(itemNumber)
	if ind > len(items) {
		fmt.Fprintf(w, "There is no item at number %v\n", ind)
		return
	}
	items[ind-1] = updatedItem
	fmt.Fprintf(w, "Item : ' %v ' Updated successfully", updatedItem)

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error : %v", err)
		return
	}

	itemNumber := r.FormValue("delete_item_number")
	ind, _ := strconv.Atoi(itemNumber)

	if ind > len(items) {
		fmt.Fprintf(w, "There is no item at number %v\n", ind)
		return
	}
	deletedItem := items[ind-1]
	items = append(items[:ind-1], items[ind:]...)
	fmt.Fprintf(w, "Item : ' %v ' deleted successfully", deletedItem)

}
