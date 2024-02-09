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

	ebt := revenue - expenses
	profit := ebt * (100 - taxRate) / 100
	ratio := ebt / profit

	formattedEbt := fmt.Sprintf("EBT: %v\n", ebt)
	formattedProfit := fmt.Sprintf("Profit: %v\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)
	fmt.Print(formattedEbt, formattedProfit, formattedRatio)
}
