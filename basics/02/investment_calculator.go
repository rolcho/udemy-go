package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationrate = 2.5
	var expectedReturnRate, investmentAmount, years float64

	fmt.Println("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationrate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
