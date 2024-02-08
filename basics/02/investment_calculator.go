package main

import (
	"fmt"
	"math"
)

func main() {
	expectedReturnRate, investmentAmount, years := 5.5, 1000.0, 10.0

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)
}
