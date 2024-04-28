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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
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
			if !isValidName {
				fmt.Println("Entered first name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("Entered email address is invalid!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
	city := "Colombo"

	switch city {
	case "Homagama", "Maharagama", "Kottawa":
		//logic goes here

	case "Piliyandala":

	case "Nugegoda":

	case "Thalawathugoda", "Malabe":

	case "Colombo":

	default:
		fmt.Print("No valid city selected")
	}
}
