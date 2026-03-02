package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const userName string = "RITU"

var items []string

func main() {

	for {
		showList()
		askUser()

	}

}

func askUser() {
	var userOption int
	fmt.Printf("\n-----------------MENU-----------------\n")
	fmt.Printf("Press 1 to add item\nPress 2 to update item\nPress 3 to delete an item\nEnter value : ")
	fmt.Scan(&userOption)
	switch userOption {
	case 1:
		addItem()
	case 2:
		updateItem()
	case 3:
		deleteItem()
	default:
		fmt.Printf("This is not a valid value! Try again\n")
	}
	fmt.Printf("---------------------------------------------\n")

}

func addItem() {

	fmt.Printf("\n-----------------ADDING ITEM-----------------")
	fmt.Printf("\nEnter the item that you want to add : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	item := scanner.Text()
	items = append(items, item)
	fmt.Printf("---------------------------------------------\n")
}

func updateItem() {
	fmt.Printf("\n-----------------UPDATING ITEM-----------------")
	fmt.Printf("\nEnter the item number that you want to update : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	itemNumber := scanner.Text()

	fmt.Printf("Enter the updated item : ")
	scanner.Scan()
	updatedItem := scanner.Text()

	ind, _ := strconv.Atoi(itemNumber)
	if ind > len(items) {
		fmt.Printf("There is no item at number %v\n", ind)
		return
	}
	items[ind-1] = updatedItem
	fmt.Printf("---------------------------------------------\n")

}

func deleteItem() {

	fmt.Printf("\n-----------------DELETING ITEM-----------------")
	fmt.Printf("\nEnter the item number that you want to delete : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	itemNumber := scanner.Text()
	ind, _ := strconv.Atoi(itemNumber)

	if ind > len(items) {
		fmt.Printf("There is no item at number %v\n", ind)
		return
	}

	items = append(items[:ind-1], items[ind:]...)
	fmt.Printf("---------------------------------------------\n")
}

func showList() {
	fmt.Printf("\n\n-----------------%v'S TO DO LIST-----------------\n", userName)
	for index, value := range items {
		var serialNumber int = 1
		serialNumber = serialNumber + index
		fmt.Printf("%d . %v\n", serialNumber, value)
	}
	fmt.Printf("-----------------------------------------------------\n")
}
