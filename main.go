package main

import (
	"strings"
	"fmt"
)

const conferenceName string = "Go Conference"
const numOfTotalTickets  int = 50
var numOfRemainingTickets int = 50
var bookings []string

func main() {
	
	greetUsers()

	for numOfRemainingTickets > 0 && len(bookings) < 50 {
		
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketAmount := validateInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
				
			bookTicket(userTickets, firstName, lastName, email)

			fmt.Println("Guests attending the conference:")
			printFirstNames()

			noTicketsRemaining := numOfRemainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first or last name that you entered is too short")
			}
			if !isValidEmail {	
				fmt.Println("The email you entered doesn't contain an @ sign")
			}
			if !isValidTicketAmount{
				fmt.Println("The number of tickets you requested is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", numOfTotalTickets, numOfRemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	for _, name := range firstNames {
		fmt.Printf("%s\n", name)
	}
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to reserve: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	numOfRemainingTickets = numOfRemainingTickets - userTickets;
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you, %v %v, we reserved %v tickets for you.\nYou'll receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", numOfRemainingTickets, conferenceName)
}