package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)

	for len(bookings) < 50 && remainingTickets > 0 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. We will send you a confirmation email shortly.\n",
				firstName, lastName, userTickets)
			fmt.Printf("We have %v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are: %v.\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
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
	fmt.Println("Welcome to the Go Conference booking application.")
}
