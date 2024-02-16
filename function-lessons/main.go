package main

import "fmt"

type transformFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	transformed := transformNumbers(&numbers, getTransformFunction(&numbers))
	fmt.Println(transformed)

	transformed = transformNumbers(&numbers, getTransformFunction(&numbers))
	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transformer transformFunction) []int {
	var transformed []int
	for _, number := range *numbers {
		transformed = append(transformed, transformer(number))
	}
	return transformed
}

func getTransformFunction(numbers *[]int) transformFunction {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}
