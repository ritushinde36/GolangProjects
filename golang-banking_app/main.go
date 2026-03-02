package main

import (
	"banking_app/fileoperations"
	"banking_app/presentation"
	"fmt"
)

var amount int

var isExit bool = false

const amountFile = "amount.txt"

func main() {
	amount_temp, err := fileoperations.GetBalance(amountFile)
	amount = amount_temp
	if err != nil {
		fmt.Printf("ERROR  -  ")
		fmt.Println(err)
		// return
		panic("Cannot proceed with the application")
	}

	// fmt.Println("WELCOME TO THE BANKING APPLICATION", randomdata.SillyName())
	fmt.Println("WELCOME TO THE BANKING APPLICATION", fileoperations.Hello)
	for {
		askUser()
		if isExit {
			break
		}
	}

}

func askUser() {
	var userChoice int
	presentation.ShowOptions()
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		checkBalance()
	case 2:
		withdrawAmount()
	case 3:
		depositMoney()
	case 4:
		exitApp()
	default:
		fmt.Printf("This is an invalid choice! Please try again.\n")
	}

}

func exitApp() {
	isExit = true
}

func checkBalance() {
	fmt.Printf("\n-----------------------------------------\n")
	fmt.Printf("This is your current balance : %v", amount)
	fmt.Printf("\n-----------------------------------------\n")

}

func withdrawAmount() {
	var withdrawAm int
	fmt.Printf("\n-----------------------------------------\n")
	fmt.Printf("Please enter the amount of money that you want to withdraw : ")
	fmt.Scan(&withdrawAm)

	checkPos := checkIfNumberIsPositive(withdrawAm)
	if !checkPos {
		fmt.Printf("This is an Invalid amount.")
		return
	}

	if withdrawAm > amount {
		fmt.Printf("You cannot withdraw more than what you have in your account. You current balance is %v\n", amount)
		fmt.Printf("\n-----------------------------------------\n")
		return
	}
	amount -= withdrawAm
	fileoperations.UpdateBalance(amount, amountFile)
	fmt.Printf("Amount %v deducted from your account\n", withdrawAm)
	checkBalance()

}

func checkIfNumberIsPositive(value int) bool {
	if value > 0 {
		return true
	} else {
		return false
	}

}

func depositMoney() {
	var depositAm int
	fmt.Printf("\n-----------------------------------------\n")
	fmt.Printf("Please enter the amount of money that you want to deposit : ")
	fmt.Scan(&depositAm)

	checkPos := checkIfNumberIsPositive(depositAm)
	if !checkPos {
		fmt.Printf("This is an Invalid amount.")
		return
	}

	amount += depositAm
	fileoperations.UpdateBalance(amount, amountFile)
	fmt.Printf("Amount %v credited from your account\n", depositAm)
	checkBalance()

}
