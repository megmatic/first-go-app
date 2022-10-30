package main

import (
	"strings"
)

func validateInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets > 0 && userTickets <= numOfRemainingTickets
	return isValidName, isValidEmail, isValidTicketAmount
}