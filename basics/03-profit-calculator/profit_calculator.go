package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	e, p, r := profitCalculator(revenue, expenses, taxRate)
	t := printProfitResult(e, p, r)

	fmt.Print(t)
}

func profitCalculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (100 - taxRate) / 100
	ratio := ebt / profit
	return ebt, profit, ratio
}

func printProfitResult(ebt, profit, ratio float64) string {
	formattedEbt := fmt.Sprintf("EBT: %v\n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %v\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)
	return formattedEbt + formattedProfit + formattedRatio
}
