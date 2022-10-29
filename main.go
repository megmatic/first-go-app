package main

import "fmt"

func main() {
	const conferenceName string = "Go Conference"
	const numOfTotalTickets  int = 50
	var numOfRemainingTickets uint = 50

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

	numOfRemainingTickets = numOfRemainingTickets - uint(userTickets);

	fmt.Printf("Thank you, %v %v, we reserved %v tickets for you.\nYou'll receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", numOfRemainingTickets, conferenceName)
}
