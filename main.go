package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(uint(userTickets), firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The following people have booked tickets: %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are out of tickets.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name must be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @ and .")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets must be between 1 and %v.\n", remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Please enter your details to book a ticket.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= uint(userTickets)

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you a confirmation email shortly.\n",
		firstName, lastName, userTickets)
	fmt.Printf("We have %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
