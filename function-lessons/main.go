package main

import "fmt"

func main() {
	s := sum(1, 5, 7, 12)

	fmt.Println(s)
}

func sum(numbers ...int) int {
	s := 0

	for _, n := range numbers {
		s += n
	}

	return s
}
