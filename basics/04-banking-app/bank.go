package main

import "fmt"

func main() {
	var balance float64 = 1000
	var isLoggedin bool = true

	for isLoggedin {
		showMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		isLoggedin = processChoice(choice, balance)
	}

}

func showMenu() {
	fmt.Println("-------------------------")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func processChoice(choice int, balance float64) bool {
	switch choice {
	case 1:
		checkBalance(balance)
	case 2:
		depositMoney(balance)
	case 3:
		withdrawMoney(balance)
	case 4:
		fmt.Println("Thank you for using GoBank!")
		return false
	default:
		fmt.Println("Invalid choice!")
	}
	return true
}

func checkBalance(balance float64) {
	fmt.Println("Your balance is:", balance)
}

func depositMoney(balance float64) {
	var amount float64
	fmt.Print("Enter the amount to deposit: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return
	}

	balance += amount
	fmt.Println("Deposited:", amount)
	fmt.Println("New balance:", balance)
}

func withdrawMoney(balance float64) {
	var amount float64
	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return
	}

	if amount > balance {
		fmt.Println("Insufficient balance!")
		return
	}

	balance -= amount
	fmt.Println("Withdrawn:", amount)
	fmt.Println("New balance:", balance)
}
