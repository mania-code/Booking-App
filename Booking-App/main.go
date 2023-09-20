package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "go conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	// var bookings = []string{}
	// var bookings []string
	bookings := []string{}

	//print type of variable
	//fmt.Printf("configureTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, conferenceTicket, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still avaiable.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for there name

		// fmt.Println(remainingTickets)
		// & is a pointer
		// fmt.Println(&remainingTickets)
		//user input
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thankyou %v %v for booking %v ticket. You will recive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets  remaning for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v ticket remaining, So you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
