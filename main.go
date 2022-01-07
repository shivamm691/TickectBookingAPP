// AUTHOUR : SHIVAM
// DATE :06:01:2022
// // LAST MODIFy : 07:01:2022
// THIS is a basic ticket booking app in go lanuage
package main

import "fmt"


func main(){



var ConferenceName = "GO Conference"
const ConferenceTickets = 50 
var remainingTickets uint = 50
fmt.Printf("Welcome to %v platfrom build in go :)", ConferenceName)
fmt.Println("we have tottel of",ConferenceTickets,  "and this many are available ",remainingTickets,)

var firstName string
var lastName string
var email string
var userTickets uint

//ask user for their name 

fmt.Printf("Enter your first name:")
fmt.Scan(&firstName)

fmt.Printf("Enter  your last name:")
fmt.Scan(&lastName)


fmt.Printf("Enter your email address: ")
fmt.Scan(&email)




fmt.Printf("hi %v %v how many ticket you want to book\n",firstName,lastName)
fmt.Scan(&userTickets)
remainingTickets = remainingTickets - userTickets

fmt.Printf("Thank you %v %v for booking %v tickest . you will  recive confirmation on email at %v\n",firstName,lastName,userTickets,email)

fmt.Printf("%v tickest remains ",remainingTickets)


}