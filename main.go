package main

import (
	"fmt"
	"strings"
)

func main() {
	conName := "Go Conference"
	const conTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conTickets is %T remainingTickets is %T, conName is %T\n", conTickets, remainingTickets, conName)

	fmt.Printf("Welcome to the %v Booking Application!\n", conName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend...")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter the First Name : ")
		fmt.Scan(&firstName)
		fmt.Print("Enter the Last Name : ")
		fmt.Scan(&lastName)
		fmt.Print("Enter the Email : ")
		fmt.Scan(&email)
		fmt.Print("Enter the Number of Tickets : ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			userTickets = 2
			fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year. Thank you!")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
