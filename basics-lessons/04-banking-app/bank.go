package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "balance.txt"

var loadedBalance float64

func main() {
	var isLoggedin bool = true

	loadedBalance = checkBalance()

	for isLoggedin {
		showMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		isLoggedin = processChoice(choice, loadedBalance)
	}

}

func storeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalance, []byte(balanceText), 0644)
}

func getBalance() (float64, error) {
	data, err := os.ReadFile(accountBalance)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("damaged data")
	}

	return balance, nil
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
		loadedBalance = checkBalance()
	case 2:
		loadedBalance = depositMoney(balance)
	case 3:
		loadedBalance = withdrawMoney(balance)
	case 4:
		fmt.Println("Thank you for using GoBank!")
		return false
	default:
		fmt.Println("Invalid choice!")
	}
	return true
}

func checkBalance() float64 {
	balance, err := getBalance()

	if err != nil {
		showError(err)
	}

	fmt.Println("Your balance is:", balance)
	return balance
}

func depositMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount to deposit: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return balance
	}

	balance += amount
	storeBalance(balance)
	fmt.Println("Deposited:", amount)
	fmt.Println("New balance:", balance)

	return balance
}

func withdrawMoney(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return balance
	}

	if amount > balance {
		fmt.Println("Insufficient balance!")
		return balance
	}

	balance -= amount
	storeBalance(balance)
	fmt.Println("Withdrawn:", amount)
	fmt.Println("New balance:", balance)
	return balance
}

func showError(err error) {
	fmt.Printf("ERROR:\n%v\n-------------\n", err)
}
