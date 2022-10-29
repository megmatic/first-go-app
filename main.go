package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const numOfTotalTickets = 50
	var numOfRemainingTickets = 50

	fmt.Printf("numOfTotalTickets is %T, numOfRemainingTickets is %T, conferenceName is %T\n", numOfTotalTickets, numOfRemainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", numOfTotalTickets, numOfRemainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	fmt.Printf("Thank you, %v, we reserved %v tickets for you.\n", firstName, userTickets)

}
