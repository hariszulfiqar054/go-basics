package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// This belong the part of the user struct and it
func (userData user) outputUserData() {
	fmt.Println(userData.firstName, userData.lastName, userData.birthdate, userData.createdAt)
}

func (userData *user) removeName() {
	userData.firstName = ""
	userData.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var userData user
	userData = user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	// Don't need to pass the userData as an argument because it's already belong to the struct and it will access
	// same struct data within that scope
	userData.outputUserData()
	userData.removeName()
	userData.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUserData(userValue user) {
	fmt.Println(userValue.firstName, userValue.lastName, userValue.birthdate, userValue.createdAt)
}
