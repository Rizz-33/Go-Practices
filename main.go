package main

import (
	"fmt"
)

var conName = "Go Conference"

const conTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
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
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application!\n", conName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend...")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var UserData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, UserData)
	fmt.Printf(("List of bookings is %v\n"), bookings)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conName)
}
