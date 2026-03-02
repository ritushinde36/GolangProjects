package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, userEmail string, userTickets int, remainingTickets int) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(userEmail, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
