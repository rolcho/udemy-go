package main

import "fmt"

func main() {
	fmt.Println("Welcome to GoBank!")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > 4 {
		fmt.Println("Invalid choice!")
		return
	}

	fmt.Println("You selected:", choice)
}
