package main

import "fmt"

func main() {
	var conName = "Go Conference"
	const conTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v Booking Application!\n", conName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend...")
}
