package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(uint(userTickets), firstName, lastName, email)
		wg.Add(1)
		go sendTicket(uint(userTickets), firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The following people have booked tickets: %v.\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are out of tickets.")
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Please enter your details to book a ticket.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: int(userTickets),
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you a confirmation email shortly.\n",
		firstName, lastName, userTickets)
	fmt.Printf("We have %v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v.\n", ticket, email)
	fmt.Println("###################################################")
	wg.Done()
}
