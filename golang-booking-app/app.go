package main

import (
	"fmt"
	"go_proj/helper"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceName string = "Go conference"
var userCount int = 0

const conferenceTickets int = 50

var remainingTickets int = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	conferenceName  string
	numberOfTickets int
	userID          int
}

func main() {

	for {

		var isValidName bool
		var isValidEmail bool
		var isValidTicketNumber bool

		firstName, lastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber = helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			greetUsers()

			bookTicket(userTickets, firstName, lastName, userEmail)
			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, userEmail)

			var firstNames []string = getFirstNames()
			fmt.Printf("These are our bookings %s\n\n\n", firstNames)

			var check_sold_out bool = remainingTickets == 0
			if check_sold_out {
				fmt.Printf("We are all sold out for the %s. Please try again next year", conferenceName)
				// fmt.Printf("\n\n\n\n\nThis is the final list of users %v", bookings)
				break
			}

		} else {

			if !isValidName {
				fmt.Println("The first name or last name that you have entered in too short! Please ensure both of them are greater than 2")
			}
			if !isValidEmail {
				fmt.Println("The Email address that you have entered in too short! Please ensure that you include the @ symbol")
			}
			if !isValidTicketNumber {
				if userTickets > remainingTickets {
					fmt.Printf("\nThe %s has only %d tickets remaining. You cannot book %d tickets. Please try again with a smaller value.\n", conferenceName, remainingTickets, userTickets)
				} else {
					fmt.Println("The number of tickets you entered is invalid")
				}
			}
		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("WELCOME TO THE %s\n\n", conferenceName)
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets int

	fmt.Printf("Please enter your first Name:-")
	fmt.Scan(&firstName)

	fmt.Printf("Please enter your last name:-")
	fmt.Scan(&lastName)

	fmt.Printf("Please enter your email address:-")
	fmt.Scan(&userEmail)

	fmt.Printf("Please enter the number of tickets you want to book:-")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets

}

func bookTicket(userTickets int, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets
	userCount = userCount + 1

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userEmail,
		numberOfTickets: userTickets,
		conferenceName:  conferenceName,
		userID:          userCount,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %s %s for booking %d tickets! You will receive a confirmation on your email - %s\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("There are a total of %d tickets, out of which, %d are available\n", conferenceTickets, remainingTickets)

}

func sendTickets(userTickets int, firstName string, lastName string, userEmail string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%d tickets for %s %s", userTickets, firstName, lastName)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Sent %v on email address %v\n", ticket, userEmail)
	fmt.Println("----------------------------------------------------------------------")
	wg.Done()
}
