package main

import "fmt"

func main() {
	conName := "Go Conference"
	const conTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conTickets is %T remainingTickets is %T, conName is %T\n", conTickets, remainingTickets, conName)

	fmt.Printf("Welcome to the %v Booking Application!\n", conName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend...")

	var userName string
	var userTickets int
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
