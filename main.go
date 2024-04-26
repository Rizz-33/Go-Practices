package main

import "fmt"

func main() {
	var conName = "Go Conference"
	const conTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conName, "Booking Application!")
	fmt.Println("We have total of", conTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend...")
}
