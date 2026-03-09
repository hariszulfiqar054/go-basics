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

	outputUserData(userData)

	// Can also be written as below with the correct order
	// userData = user{
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }

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
