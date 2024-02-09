package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	e, p, r := calculateFinance(revenue, expenses, taxRate)
	t := textFinance(e, p, r)

	fmt.Print(t)
}

func getUserInput(infoText string) (value float64) {
	fmt.Print(infoText, ": ")
	fmt.Scan(&value)
	return
}

func calculateFinance(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (100 - taxRate) / 100
	ratio = ebt / profit
	return
}

func textFinance(ebt, profit, ratio float64) (formattedResult string) {
	formattedEbt := fmt.Sprintf("EBT: %v\n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %v\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)
	formattedResult = formattedEbt + formattedProfit + formattedRatio
	return
}
