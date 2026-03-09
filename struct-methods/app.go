package main

import (
	"fmt"
	"struct.com/user/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var userData user.User
	userData = *user.New(firstName, lastName, birthdate)
	// Don't need to pass the userData as an argument because it's already belong to the struct and it will access
	// same struct data within that scope
	userData.OutputUserData()
	userData.RemoveName()
	userData.OutputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
