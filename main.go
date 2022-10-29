package main

import (
	"strings"
	"fmt"
)

func main() {
	const conferenceName string = "Go Conference"
	const numOfTotalTickets  int = 50
	var numOfRemainingTickets int = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", numOfTotalTickets, numOfRemainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		if  userTickets > numOfRemainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", numOfRemainingTickets, userTickets)
			continue
		}
		numOfRemainingTickets = numOfRemainingTickets - userTickets;
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you, %v %v, we reserved %v tickets for you.\nYou'll receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", numOfRemainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		noTicketsRemaining := numOfRemainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}
