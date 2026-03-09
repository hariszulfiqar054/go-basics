package admin

import (
	"struct.com/struct-app/user"
)

// Its like an inheritance it will also have access to user properties and also the methods
type Admin struct {
	email    string
	password string
	User     user.User
}

func New(email, password, firstName, lastName, birthdate string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User:     *user.New(firstName, lastName, birthdate),
	}
}
