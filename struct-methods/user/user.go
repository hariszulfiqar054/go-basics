package user

import (
	"fmt"
	"time"
)

// In order to export struct properties and methods they needed to be capitalized,
// otherwise they will be unexported and only accessible within the package.
type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

// This belong the part of the user struct and it
func (userData *User) OutputUserData() {
	fmt.Println(userData.FirstName, userData.LastName, userData.Birthdate, userData.CreatedAt)
}

func (userData *User) RemoveName() {
	userData.FirstName = ""
	userData.LastName = ""
}

func New(firstName, lastName, birthdate string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
