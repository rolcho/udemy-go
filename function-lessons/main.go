package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	s := sum(10, 1, 5, 7, 12)
	fmt.Println(s)

	anotherSum := sum(1, numbers...)
	fmt.Println(anotherSum)
}

func sum(startingValue int, numbers ...int) int {
	s := startingValue

	for _, n := range numbers {
		s += n
	}

	return s
}
