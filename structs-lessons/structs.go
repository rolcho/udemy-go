package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// reciever argument passing value
func (u user) outputUserDetail() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// reciever argument passing pointer
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function: conventional naming newType
func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(userFirstName, userLastName, userBirthDate)

	// ... do something awesome with that gathered data!
	appUser.outputUserDetail()
	appUser.clearUserName()
	appUser.outputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
