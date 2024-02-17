package main

import "fmt"

func main() {

	fact := factorial(6)
	fmt.Println(fact)
}

func factorial(n int) int {
	if n > 1 {
		return n * factorial(n-1)
	}
	return 1
}
