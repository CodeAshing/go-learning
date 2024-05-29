package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var conferenceAvailableTicket uint = 50

	fmt.Printf("Welcome to %s\n", conferenceName)
	fmt.Printf("We have %d tickets available out of %d\n", conferenceAvailableTicket, conferenceTicket)
	fmt.Println ("buy your ticket now")	 

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTicket uint
	booking := []string{}

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter the number of tickets you want to buy: ")
	fmt.Scan(&userTicket)

	booking = append(booking, userFirstName + " " + userLastName)

	conferenceAvailableTicket = conferenceTicket - userTicket

	fmt.Printf("Thank you %s %s, you have successfully bought %d tickets to %s you will get confirmation at %v\n", userFirstName, userLastName, userTicket, conferenceName, userEmail)
	fmt.Printf("We have %d tickets available out of %d\n", conferenceAvailableTicket, conferenceTicket)

	fmt.Printf("These are all our booking: %v\n", booking)


}