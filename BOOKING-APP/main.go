package main
import (
	"booking-app/slider"
	"fmt"
	"sync"
	"time"
)

const eventTickets = 100

var eventName = "MEGA EVENT"
var remainingTickets uint = 100
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//bookings := []string{}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := slider.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our event is booked out. Come back next year.")

		}

	} else {
		if !isValidName {
			fmt.Println("First name or Last name you enterd is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", eventTickets, remainingTickets)
	fmt.Println("So HURRY UP !!! Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)

	var email string
	fmt.Println("Enter your E-mail address :")
	fmt.Scan(&email)

	var userTickets uint
	fmt.Println("Enter number of tickets:") // have to be careful for this line.()
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}
