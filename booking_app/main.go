package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]map[string]string, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// const conferenceTickets int = 50
	// var remainingTickets uint = 50
	// conferenceName := "Go Conference"
	// bookings := []string{}

	// fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint

	// // asking for user input
	// fmt.Println("Enter Your First Name: ")
	// fmt.Scanln(&firstName)

	// fmt.Println("Enter Your Last Name: ")
	// fmt.Scanln(&lastName)

	// fmt.Println("Enter Your Email: ")
	// fmt.Scanln(&email)

	// fmt.Println("Enter number of tickets: ")
	// fmt.Scanln(&userTickets)

	// // book ticket in system
	// remainingTickets = remainingTickets - userTickets
	// bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// fmt.Printf("These are all our bookings: %v\n", bookings)

	// const conferenceTickets int = 50
	// var remainingTickets uint = 50
	// conferenceName := "Go Conference"
	// bookings := []string{}

	// fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	// for {
	// 	var firstName string
	// 	var lastName string
	// 	var email string
	// 	var userTickets uint

	// 	// asking for user input
	// 	fmt.Println("Enter Your First Name: ")
	// 	fmt.Scanln(&firstName)

	// 	fmt.Println("Enter Your Last Name: ")
	// 	fmt.Scanln(&lastName)

	// 	fmt.Println("Enter Your Email: ")
	// 	fmt.Scanln(&email)

	// 	fmt.Println("Enter number of tickets: ")
	// 	fmt.Scanln(&userTickets)

	// 	// validate user input
	// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// 	isValidEmail := strings.Contains(email, "@")
	// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// 	if isValidName && isValidEmail && isValidTicketNumber {

	// 		// book ticket in system
	// 		remainingTickets = remainingTickets - userTickets
	// 		bookings = append(bookings, firstName+" "+lastName)

	// 		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// 		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// 		// print only first names
	// 		firstNames := []string{}
	// 		for _, booking := range bookings {
	// 			var names = strings.Fields(booking)
	// 			firstNames = append(firstNames, names[0])
	// 		}
	// 		fmt.Printf("The first names %v\n", firstNames)

	// 		// exit application if no tickets are left
	// 		if remainingTickets == 0 {
	// 			// end program
	// 			fmt.Println("Our conference is booked out. Come back next year.")
	// 			break
	// 		}
	// 	} else {
	// 		if !isValidName {
	// 			fmt.Println("first name or last name you entered is too short")
	// 		}
	// 		if !isValidEmail {
	// 			fmt.Println("email address you entered doesn't contain @ sign")
	// 		}
	// 		if !isValidTicketNumber {
	// 			fmt.Println("number of tickets you entered is invalid")
	// 		}
	// 		continue
	// 	}
	// }

	// // Switch statement example
	// city := "London"

	// switch city {
	// case "New York":
	// 	// booking for New York conference
	// case "Singapore", "Hong Kong":
	// 	// booking for Singapore & Hong Kong conference
	// case "London", "Berlin":
	// 	// booking for London & Berlin conference
	// case "Mexico City":
	// 	// booking for Mexico City conference
	// default:
	// 	fmt.Print("No valid city selected")
	// }

	// greetUsers()

	// for {

	// 	firstName, lastName, email, userTickets := getUserInput()
	// 	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	// 	if isValidName && isValidEmail && isValidTicketNumber {

	// 		bookTicket(userTickets, firstName, lastName, email)

	// 		firstNames := printFirstNames()
	// 		fmt.Printf("The first names %v\n", firstNames)

	// 		if remainingTickets == 0 {
	// 			// end program
	// 			break
	// 		}
	// 	} else {
	// 		if !isValidName {
	// 			fmt.Println("firt name or last name you entered is too short")
	// 		}
	// 		if !isValidEmail {
	// 			fmt.Println("email address you entered doesn't contain @ sign")
	// 		}
	// 		if !isValidTicketNumber {
	// 			fmt.Println("number of tickets you entered is invalid")
	// 		}
	// 		continue
	// 	}
	// }

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames() []string {
	firstNames := []string{}

	// for _, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, firstName+" "+lastName)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets = remainingTickets - userTickets

// 	// create user map
// 	var user = make(map[string]string)
// 	user["firstName"] = firstName
// 	user["lastName"] = lastName
// 	user["email"] = email
// 	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

// 	bookings = append(bookings, user)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create user map
	// var user = User{
	// 	firstName:       firstName,
	// 	lastName:        lastName,
	// 	email:           email,
	// 	numberOfTickets: userTickets,
	// }

	// bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
