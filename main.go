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

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	bookings := []string{}

	fmt.Print("Enter the First Name : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter the Last Name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter the Email : ")
	fmt.Scan(&email)
	fmt.Print("Enter the Number of Tickets : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\nThe whole slice : %v\n", bookings)
	fmt.Printf("The first value : %v\n", bookings[0])
	fmt.Printf("Array type : %T\n", bookings)
	fmt.Printf("Array length : %v\n", len(bookings))

	userTickets = 2
	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conName)
}
