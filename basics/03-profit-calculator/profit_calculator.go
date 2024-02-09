package main

import (
	"fmt"
)

func main() {

	revenue := inputValue("Revenue")
	expenses := inputValue("Expenses")
	taxRate := inputValue("Tax Rate")

	e, p, r := profitCalculator(revenue, expenses, taxRate)
	t := printProfitResult(e, p, r)

	fmt.Print(t)
}

func inputValue(message string) (value float64) {
	fmt.Print(message, ": ")
	fmt.Scan(&value)
	return
}

func profitCalculator(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (100 - taxRate) / 100
	ratio = ebt / profit
	return
}

func printProfitResult(ebt, profit, ratio float64) (formattedResult string) {
	formattedEbt := fmt.Sprintf("EBT: %v\n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %v\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)
	formattedResult = formattedEbt + formattedProfit + formattedRatio
	return
}
