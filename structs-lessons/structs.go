package main

import (
	"fmt"

	"example.com/structs/user"
)

// custom type
type str string

// add method tocustom type
func (text str) log() {
	fmt.Println(text)
}

func main() {
	//basic usage of custom type & method
	var myName str = "Hi"
	myName.log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return

	}

	admin := user.NewAdmin("a@b.c", "Admin", "admin.com", userBirthDate)

	admin.OutputUserDetail()

	appUser.OutputUserDetail()
	appUser.ClearUserName()
	appUser.OutputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
