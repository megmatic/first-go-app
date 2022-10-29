package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const numOfTotalTickets = 50
	var numOfRemainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", numOfTotalTickets, "tickets and", numOfRemainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}
